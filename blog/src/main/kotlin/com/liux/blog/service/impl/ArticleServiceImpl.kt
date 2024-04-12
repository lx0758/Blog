package com.liux.blog.service.impl

import com.github.pagehelper.Page
import com.github.pagehelper.PageHelper
import com.liux.blog.bean.event.BaseInfoUpdateEvent
import com.liux.blog.bean.po.Article
import com.liux.blog.bean.po.ArticleUrl
import com.liux.blog.bean.po.ArticleUrlStatusEnum
import com.liux.blog.bean.po.Tag
import com.liux.blog.dao.ArticleMapper
import com.liux.blog.dao.ArticleUrlMapper
import com.liux.blog.dao.TagMapper
import com.liux.blog.service.ArticleService
import com.liux.blog.service.FragmentService
import com.liux.blog.util.PageHelperUtil
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.context.ApplicationEventPublisher
import org.springframework.stereotype.Service
import org.springframework.transaction.annotation.Transactional
import java.util.*

@Service
class ArticleServiceImpl: ArticleService {

    @Autowired
    private lateinit var applicationEventPublisher: ApplicationEventPublisher

    @Autowired
    private lateinit var articleMapper: ArticleMapper
    @Autowired
    private lateinit var articleUrlMapper: ArticleUrlMapper
    @Autowired
    private lateinit var tagMapper: TagMapper

    @Autowired
    private lateinit var fragmentService: FragmentService

    override fun listByPage(pageNum: Int): Page<Article> {
        val page = PageHelper.startPage<Article>(pageNum, 10)
        articleMapper.selectByPage()
        return page
    }

    override fun listByArchive(pageNum: Int): Page<Article> {
        val page = PageHelper.startPage<Article>(pageNum, 10)
        articleMapper.selectByArchive()
        return page
    }

    override fun listByCategory(categoryId: Int, pageNum: Int): Page<Article> {
        val page = PageHelper.startPage<Article>(pageNum, 10)
        articleMapper.selectByCategory(categoryId)
        return page
    }

    override fun listByTag(tagId: Int, pageNum: Int): Page<Article> {
        val page = PageHelper.startPage<Article>(pageNum, 10)
        articleMapper.selectByTag(tagId)
        return page
    }

    override fun listBySearch(): List<Article> {
        return articleMapper.selectBySearch()
    }

    override fun listBySitemap(): List<Article> {
        return articleMapper.selectBySitemap()
    }

    override fun listByExport(): List<Article> {
        return articleMapper.selectByExport()
    }

    override fun listByAdminDashboard(pageNum: Int, pageSize: Int): Page<Article> {
        val page = PageHelper.startPage<Article>(pageNum, pageSize)
        articleMapper.selectByAdminDashboard()
        return page
    }

    override fun listByAdmin(
        title: String?,
        categoryId: Int?,
        url: String?,
        commentStatus: Int?,
        status: Int?,
        pageNum: Int,
        pageSize: Int,
        orderName: String?,
        orderMethod: String?,
    ): Page<Article> {
        PageHelperUtil.orderBy(PageHelperUtil.Type.ARTICLE, orderName, orderMethod)
        val page = PageHelper.startPage<Article>(pageNum, pageSize)
        articleMapper.selectByAdmin(Article(
            id = 0,
            categoryId = categoryId,
            title = title,
            commentStatus = commentStatus,
            status = status,
        ), url)
        return page
    }

    override fun getByBlog(id: Int): Article? {
        return articleMapper.getByBlog(id)?.apply {
            views = views?.plus(1)
            articleMapper.updateByPrimaryKeySelective(Article(this.id, views = this.views))
            fragmentService.apply(this)
        }
    }

    override fun getByPrev(articleId: Int): Article? {
        return articleMapper.getByPrev(articleId)
    }

    override fun getByNext(articleId: Int): Article? {
        return articleMapper.getByNext(articleId)
    }

    override fun getByAdmin(id: Int): Article? {
        return articleMapper.getByAdmin(id)
    }

    override fun getCountByDashboard(): Int {
        return articleMapper.getCount()
    }

    override fun getViewsByDashboard(): Int {
        return articleMapper.getViews()
    }

    @Transactional(rollbackFor = [Exception::class])
    override fun addByAdmin(
        userId: Int,
        title: String,
        categoryId: Int,
        content: String,
        url: String?,
        weight: Int?,
        commentStatus: Int,
        status: Int,
        tags: Array<String>?
    ): Article {
        val article = Article(
            title = title,
            content = content,
            categoryId = categoryId,
            authorId = userId,
            weight = weight,
            views = 0,
            commentStatus = commentStatus,
            status = status,
            createTime = Date(),
            updateTime = null,
        )
        articleMapper.insertSelective(article)

        if (!url.isNullOrBlank()) {
            val articleUrl = ArticleUrl(
                url = url,
                articleId = article.id,
                status = ArticleUrlStatusEnum.CURRENT.value,
                createTime = article.createTime
            )
            articleUrlMapper.insertSelective(articleUrl)
        }

        linkArticleAndTags(article.id!!, tags)

        applicationEventPublisher.publishEvent(BaseInfoUpdateEvent.ARTICLE)

        return article
    }

    @Transactional(rollbackFor = [Exception::class])
    override fun updateByAdmin(
        id: Int,
        title: String,
        categoryId: Int,
        content: String,
        url: String?,
        weight: Int?,
        commentStatus: Int,
        status: Int,
        tags: Array<String>?
    ): Int {
        val article = Article(
            id = id,
            title = title,
            content = content,
            categoryId = categoryId,
            weight = weight,
            commentStatus = commentStatus,
            status = status,
            updateTime = Date(),
        )
        val updateRow = articleMapper.updateByPrimaryKeyNullable(article)
        if (updateRow < 1) return updateRow

        if (!url.isNullOrBlank()) {
            articleUrlMapper.updateToHistoryByArticleId(id)
            if (articleUrlMapper.getByPrimaryKey(url) != null) {
                articleUrlMapper.deleteByPrimaryKey(url)
            }
            val articleUrl = ArticleUrl(
                url = url,
                articleId = id,
                status = ArticleUrlStatusEnum.CURRENT.value,
                createTime = article.updateTime,
            )
            articleUrlMapper.insertSelective(articleUrl)
        }

        val tagPOs = tagMapper.selectByArticleId(id)
        val tagStrings = ArrayList<String>().apply {
            if (tags != null) addAll(tags)
        }
        // 先检查已存在的 tag 和新的 tag 有没有一样的
        tagPOs.forEach {
            val tagId = it.id!!
            val tagName = it.name!!
            if (tagStrings.contains(tagName)) {
                // 如果一样则不处理, 从新数组中移除避免后续检查
                tagStrings.remove(tagName)
            } else {
                // 如果不一样则移除关联
                // 并检查是否还被别的文章关联, 没有则删除
                tagMapper.deleteByRemoveArticleLink(id, tagId)
                if (tagMapper.getByTagCount(tagId) == 0) {
                    tagMapper.deleteByPrimaryKey(tagId)
                }
            }
        }
        // 上面处理剩下的都是新的 tag
        // 只需要 查找/创建 然后关联起来即可
        linkArticleAndTags(id, tagStrings.toTypedArray())

        applicationEventPublisher.publishEvent(BaseInfoUpdateEvent.ARTICLE)

        return updateRow
    }

    @Transactional(rollbackFor = [Exception::class])
    override fun updateStatusByAdmin(id: Int, status: Int): Int {
        val updateRow = articleMapper.updateByPrimaryKeySelective(Article(
            id = id,
            status = status,
            updateTime = Date(),
        ))
        if (updateRow < 1) return updateRow

        applicationEventPublisher.publishEvent(BaseInfoUpdateEvent.ARTICLE)

        return updateRow
    }

    @Transactional(rollbackFor = [Exception::class])
    override fun deleteByAdmin(id: Int): Int {
        val deleteRow = articleMapper.deleteByPrimaryKey(id)
        if (deleteRow < 1) return deleteRow

        articleUrlMapper.deleteByArticleId(id)

        val tagPOs = tagMapper.selectByArticleId(id)
        tagPOs.forEach {
            val tagId = it.id!!
            tagMapper.deleteByRemoveArticleLink(id, tagId)
            if (tagMapper.getByTagCount(tagId) == 0) {
                tagMapper.deleteByPrimaryKey(tagId)
            }
        }

        applicationEventPublisher.publishEvent(BaseInfoUpdateEvent.ARTICLE)

        return deleteRow
    }

    private fun linkArticleAndTags(id: Int, tags: Array<String>?) {
        tags?.forEach {
            var tag = tagMapper.getByName(it)
            if (tag == null) {
                tag = Tag(
                    name = it,
                    createTime = Date(),
                )
                tagMapper.insertSelective(tag)
            }
            tagMapper.insertByAddArticleLink(id, tag.id!!)
        }
    }
}
