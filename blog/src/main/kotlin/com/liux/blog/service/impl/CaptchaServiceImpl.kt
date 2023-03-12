package com.liux.blog.service.impl

import com.liux.blog.service.CaptchaService
import jakarta.servlet.http.HttpSession
import org.springframework.stereotype.Service
import java.awt.Color
import java.awt.Font
import java.awt.image.BufferedImage
import java.io.ByteArrayOutputStream
import java.security.SecureRandom
import java.time.LocalDateTime
import javax.imageio.ImageIO

@Service
class CaptchaServiceImpl : CaptchaService {

    companion object {
        private const val CAPTCHA_TEXT_LENGTH = 4
        private const val CAPTCHA_IMAGE_WIDTH = 150
        private const val CAPTCHA_IMAGE_HEIGHT = 50
        private const val SESSION_LAST_CAPTCHA = "session_config"
        private const val SESSION_LAST_DATETIME = "session_datetime"
        private val CODE_SEQUENCE = charArrayOf(
            'A', 'B', 'C', 'D', 'E', 'F', 'G',
            'H', 'I', 'J', 'K', 'L', 'M', 'N',
            'P', 'Q', 'R', 'S', 'T',
            'U', 'V', 'W', 'X', 'Y', 'Z',
            '2', '3', '4', '5', '6', '7', '8', '9'
        )
        private val RANDOM = SecureRandom()
    }

    override fun generate(session: HttpSession, type: Int): ByteArray {
        val text = generateRandomText(CAPTCHA_TEXT_LENGTH)
        val image = generateCaptchaImage(CAPTCHA_IMAGE_WIDTH, CAPTCHA_IMAGE_HEIGHT, text)
        session.setAttribute(SESSION_LAST_CAPTCHA, text)
        session.setAttribute(SESSION_LAST_DATETIME, LocalDateTime.now())
        return image
    }

    override fun verify(session: HttpSession, captcha: String, type: Int, validPeriodMinute: Int): Boolean {
        val sessionCaptcha = session.getAttribute(SESSION_LAST_CAPTCHA) ?: return false
        if (sessionCaptcha != captcha.uppercase()) return false
        if (validPeriodMinute > 0) {
            val sessionLastDatetime = session.getAttribute(SESSION_LAST_DATETIME) ?: return false
            if (sessionLastDatetime !is LocalDateTime) return false
            if (!sessionLastDatetime.plusMinutes(validPeriodMinute.toLong()).isAfter(LocalDateTime.now())) return false
        }
        session.removeAttribute(SESSION_LAST_CAPTCHA)
        session.removeAttribute(SESSION_LAST_DATETIME)
        return true
    }

    fun generateCaptchaImage(
        width: Int, height: Int, text: String
    ): ByteArray {
        val count = text.length
        val bufferedImage = BufferedImage(width, height, BufferedImage.TYPE_INT_RGB)
        val graphics = bufferedImage.createGraphics()
        // 设置字体
        graphics.font = Font("Fixedsys", Font.PLAIN, height - 5)
        // 画背景
        graphics.color = Color.WHITE
        graphics.fillRect(0, 0, width, height)
        // 画干扰线
        for (i in 0 until count * 3) {
            graphics.color = generateRandomColor()
            graphics.drawLine(
                RANDOM.nextInt(width),
                RANDOM.nextInt(height),
                RANDOM.nextInt(width),
                RANDOM.nextInt(height)
            )
        }
        // 画噪点
        for (i in 0 until count * 5) {
            graphics.color = generateRandomColor()
            graphics.fillRect(
                RANDOM.nextInt(width),
                RANDOM.nextInt(height),
                4,
                4
            )
        }
        // 画字符
        val charWidth = width / count
        val charHeight = height - 5
        for (i in 0 until count) {
            graphics.color = generateRandomColor()
            val char = text[i].toString()
            val x = i * charWidth
            val degree = RANDOM.nextInt() % 40
            graphics.rotate(degree * Math.PI / 180, x.toDouble(), height * 0.5)
            graphics.drawString(char, x, charHeight)
            graphics.rotate(-degree * Math.PI / 180, x.toDouble(), height * 0.5)
        }
        // 获得数据
        val bufferOutputStream = ByteArrayOutputStream()
        ImageIO.write(bufferedImage, "JPEG", bufferOutputStream)
        return bufferOutputStream.toByteArray()
    }

    private fun generateRandomText(length: Int): String {
        val text = StringBuffer()
        for (i in 0 until length) {
            val strRandom = CODE_SEQUENCE[RANDOM.nextInt(CODE_SEQUENCE.size)].toString()
            text.append(strRandom)
        }
        return text.toString()
    }

    private fun generateRandomColor(): Color {
        return Color(RANDOM.nextInt(256), RANDOM.nextInt(256), RANDOM.nextInt(256))
    }
}