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
            "博客邮件发送测试",
"""
来自博客的自动邮件：
Automated emails from blogs:

如果你能看到这封邮件，说明博客的邮件参数已正确配置，祝您使用愉快！
If you can see this email, it means that the email parameters of the blog are correctly configured, I wish you a happy use!
""".trimIndent(),
        )
    }

    override fun sendCommentAddEmail(
        articleId: Int,
        nickname: String,
        content: String,
    ) {
        val ownerEmail = themeService.getCacheBlogInfo().ownerEmail
        checkAndSendEmail(ownerEmail, null, "评论通知") {
            val article = articleService.getByAdmin(articleId)!!
            val url = "https://${themeService.getCacheBlogInfo().siteDomain}/article/${if (!article.url.isNullOrBlank()) article.url else articleId}"
"""
${themeService.getCacheBlogInfo().ownerNickname}:

您的文章《${article.title}》收到了来自 $nickname 的评论。内容如下：
Your article "${article.title}" received a comment from "$nickname". The contents are as follows:

$content

文章链接如下：
The article is linked below:
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
        val ownerEmail = themeService.getCacheBlogInfo().ownerEmail
        checkAndSendEmail(comment.email!!, ownerEmail, "评论收到新的回复") {
            val article = articleService.getByAdmin(articleId)!!
            val url = "https://${themeService.getCacheBlogInfo().siteDomain}/article/${if (!article.url.isNullOrBlank()) article.url else articleId}"
            val owner = themeService.getCacheBlogInfo().ownerNickname
"""
${comment.author!!}:

您在文章《${article.title}》发出的评论收到了来自 $owner 的回复。内容如下：
Your comment on the "${article.title}" received a reply from the "$owner". The contents are as follows:

$content

文章链接如下：
The article is linked below:
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