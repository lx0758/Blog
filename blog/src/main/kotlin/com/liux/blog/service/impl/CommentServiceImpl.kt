package com.liux.blog.service.impl

import com.github.pagehelper.Page
import com.github.pagehelper.PageHelper
import com.liux.blog.bean.po.Comment
import com.liux.blog.bean.po.STATE_NOT_ACTIVATED
import com.liux.blog.dao.CommentMapper
import com.liux.blog.service.CommentService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service
import java.util.*

@Service
class CommentServiceImpl: CommentService {

    @Autowired
    private lateinit var commentMapper: CommentMapper

    override fun listByPage(pageNum: Int, pageSize: Int): Page<Comment> {
        val page = PageHelper.startPage<Comment>(pageNum, pageSize)
        commentMapper.selectByPage()
        return page
    }

    override fun listByArticle(articleId: Int, pageNum: Int): Page<Comment> {
        val page = PageHelper.startPage<Comment>(pageNum, 10)
        commentMapper.selectByArticleId(articleId)
        return page
    }

    override fun getCommentById(id: Int): Comment? {
        return commentMapper.selectByPrimaryKey(id)
    }

    override fun addByBlog(
        articleId: Int,
        parentId: Int?,
        nickname: String?,
        email: String?,
        url: String?,
        content: String,
        ip: String?,
        ua: String?
    ) {
        val comment = Comment(
            id = 0,
            articleId = articleId,
            parentId = parentId,
            author = nickname,
            authorId = null,
            email = email,
            url = url,
            content = content,
            ip = ip,
            ua = ua,
            status = STATE_NOT_ACTIVATED,
            createTime = Date(),
            updateTime = null,
        )
        commentMapper.insert(comment)
    }
}
