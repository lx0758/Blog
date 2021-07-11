package com.liux.blog.service.impl

import com.liux.blog.bean.po.Features
import com.liux.blog.service.*
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.mail.javamail.JavaMailSenderImpl
import org.springframework.mail.javamail.MimeMessageHelper
import org.springframework.stereotype.Service
import java.util.*
import javax.net.SocketFactory
import javax.net.ssl.SSLSocketFactory

@Service
class EmailServiceImpl : EmailService {

    @Autowired
    private lateinit var articleService: ArticleService

    @Autowired
    private lateinit var commentService: CommentService

    @Autowired
    private lateinit var featuresService: FeaturesService

    @Autowired
    private lateinit var themeService: ThemeService

    override fun testEmail(to: String) {
        val smtp = featuresService.readSMTP() ?: throw NullPointerException("SMTP no configuration")
        if (smtp.enable != true) throw NullPointerException("SMTP is disable")
        sendEmailImpl(
            smtp,
            to,
            null,
            "测试邮件",
"""
你好：
    如果你能看到这封邮件，说明博客的邮件参数已正确配置，祝您使用愉快！
""".trimIndent(),
        )
    }

    override fun sendCommentAddEmail(
        articleId: Int,
        nickname: String,
        content: String,
    ) {
        val ownerEmail = themeService.getCacheBase().ownerEmail
        checkAndSendEmail(ownerEmail, null, "评论通知") {
            val article = articleService.getByAdmin(articleId)!!
            val url = "https://" + themeService.getCacheBase().siteDomain + "/article/" + articleId
"""
你好, ${themeService.getCacheBase().ownerNickname}:
    您的文章《${article.title}》收到了来自 $nickname 的评论。内容如下：

$content

$url
""".trimIndent()
        }
    }

    override fun sendCommentReplayEmail(
        articleId: Int,
        commentId: Int,
        content: String,
    ) {
        val comment = commentService.getCommentById(commentId)!!
        val ownerEmail = themeService.getCacheBase().ownerEmail
        checkAndSendEmail(comment.email!!, ownerEmail, "评论收到新的回复") {
            val article = articleService.getByAdmin(articleId)!!
            val url = "https://" + themeService.getCacheBase().siteDomain + "/article/" + articleId
            val owner = themeService.getCacheBase().ownerNickname
"""
你好，${comment.author!!}:
    您在文章《${article.title}》发出的评论收到了来自 $owner 的回复。内容如下：

$content

$url
""".trimIndent()
        }
    }

    private fun checkAndSendEmail(to: String, cc: String?, subject: String, function: () -> String) {
        val smtp = featuresService.readSMTP() ?: return
        if (smtp.enable != true) return
        sendEmailImpl(smtp, to, cc, subject, function())
    }

    fun sendEmailImpl(smtp: Features.SMTP, to: String, cc: String?, subject: String, content: String) {
        val javaMailSenderImpl = JavaMailSenderImpl().apply {
            host = smtp.hostname
            port = smtp.port!!
            username = smtp.username
            password = smtp.password
            javaMailProperties = Properties().apply {
                setProperty("mail.smtp.timeout", (10 * 1000).toString())
                setProperty("mail.smtp.auth", "false")
                if (smtp.ssl == true) {
                    setProperty("mail.smtp.socketFactory.class", SSLSocketFactory::class.java.canonicalName)
                } else {
                    setProperty("mail.smtp.socketFactory.class", SocketFactory::class.java.canonicalName)
                }
            }
        }

        val helper = MimeMessageHelper(javaMailSenderImpl.createMimeMessage()).apply {
            setFrom(smtp.fromEmail!!, smtp.fromName ?: smtp.fromEmail)
            setTo(to)
            if (cc != null) setCc(cc)
            setSentDate(Date())
            setSubject(subject)
            setText(content)
        }

        javaMailSenderImpl.send(helper.mimeMessage)
    }
}