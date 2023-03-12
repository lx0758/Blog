package com.liux.blog.service.impl

import com.github.pagehelper.Page
import com.github.pagehelper.PageHelper
import com.liux.blog.bean.po.Comment
import com.liux.blog.bean.po.CommentStatusEnum
import com.liux.blog.dao.CommentMapper
import com.liux.blog.service.CommentService
import com.liux.blog.service.EmailService
import com.liux.blog.service.LocationService
import com.liux.blog.util.PageHelperUtil
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service
import org.springframework.transaction.annotation.Transactional
import java.util.*

@Service
class CommentServiceImpl: CommentService {

    @Autowired
    private lateinit var commentMapper: CommentMapper
    @Autowired
    private lateinit var emailService: EmailService
    @Autowired
    private lateinit var locationService: LocationService

    override fun listByArticle(articleId: Int, pageNum: Int): Page<Comment> {
        val page = PageHelper.startPage<Comment>(pageNum, 10)
        commentMapper.selectByArticleId(articleId)
        return page
    }

    override fun listByAdmin(
        articleId: Int?,
        author: String?,
        email: String?,
        ip: String?,
        content: String?,
        status: Int?,
        pageNum: Int,
        pageSize: Int,
        orderName: String?,
        orderMethod: String?,
    ): Page<Comment> {
        PageHelperUtil.orderBy(PageHelperUtil.Type.COMMENT, orderName, orderMethod)
        val page = PageHelper.startPage<Comment>(pageNum, pageSize)
        commentMapper.selectByAdmin(Comment(
            articleId = articleId,
            author = author,
            email = email,
            ip = ip,
            content = content,
            status = status,
        ))
        return page
    }

    override fun getCommentById(id: Int): Comment? {
        return commentMapper.getByPrimaryKey(id)
    }

    override fun getCountByDashboard(): Int {
        return commentMapper.getCount()
    }

    override fun addByBlog(
        articleId: Int,
        parentId: Int?,
        nickname: String,
        email: String,
        url: String?,
        content: String,
        ip: String,
        ua: String,
    ) {
        val finalParentId = findTopParentId(parentId)
        val comment = Comment(
            articleId = articleId,
            parentId = finalParentId,
            author = nickname,
            authorId = null,
            email = email,
            url = url,
            content = content,
            ip = ip,
            location = locationService.getLocationFromIp(ip),
            ua = ua,
            status = CommentStatusEnum.PENDING.value,
            createTime = Date(),
            updateTime = null,
        )
        commentMapper.insert(comment)

        emailService.sendCommentAddEmail(articleId, nickname, content)
    }

    override fun addByAdmin(
        userId: Int,
        nickname: String,
        email: String,
        articleId: Int,
        parentId: Int,
        content: String,
        ip: String,
        ua: String,
        emailNotify: Boolean,
    ) {
        val finalParentId = findTopParentId(parentId)
        val comment = Comment(
            articleId = articleId,
            parentId = finalParentId,
            author = nickname,
            authorId = null,
            email = email,
            url = null,
            content = content,
            ip = ip,
            location = locationService.getLocationFromIp(ip),
            ua = ua,
            status = CommentStatusEnum.AUDITED.value,
            createTime = Date(),
            updateTime = null,
        )
        commentMapper.insert(comment)

        if (emailNotify) {
            emailService.sendCommentReplayEmail(articleId, parentId, content)
        }
    }

    override fun updateByAdmin(id: Int, status: Int): Int {
        val comment = Comment(
            id = id,
            status = status,
            updateTime = Date(),
        )
        return commentMapper.updateByPrimaryKeySelective(comment)
    }

    @Transactional(rollbackFor = [Exception::class])
    override fun deleteByAdmin(id: Int): Int {
        val deleteRow = commentMapper.deleteByPrimaryKey(id)
        if (deleteRow > 0) {
            commentMapper.deleteByCleanChild(id)
        }
        return deleteRow
    }

    /**
     * 找到最顶层的评论ID
     */
    private fun findTopParentId(parentId: Int?): Int? {
        if (parentId == null) return null
        var finalParentId = parentId
        while (true) {
            val localComment = commentMapper.getByPrimaryKey(finalParentId!!) ?: break
            val localParentId = localComment.parentId ?: break
            finalParentId = localParentId
        }
        return finalParentId
    }
}
