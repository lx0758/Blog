package com.liux.blog.bean.vo

import com.liux.blog.bean.po.Config
import com.liux.blog.bean.po.Link
import com.liux.blog.bean.po.User
import java.util.*

data class BlogVO(
    var siteDomain: String,
    var siteTitle: String,
    var siteDescription: String,
    var siteKeywords: String,

    var siteBeianIcp: String,
    var siteBeianPs: String,
    var siteBaidu: String,
    var siteCreateYear: String,

    var siteArticleCount: Int,
    var siteCategoryCount: Int,
    var siteTagCount: Int,

    var ownerAvatar: String,
    var ownerNickname: String,
    var ownerDescription: String,
    var ownerEmail: String,

    var ownerGithub: String,
    var ownerWeibo: String,
    var ownerGoogle: String,
    var ownerTwitter: String,
    var ownerFacebook: String,
    var ownerStackOverflow: String,
    var ownerYoutube: String,
    var ownerInstagram: String,
    var ownerSkype: String,

    var links: List<LinkVO> = emptyList(),

    var updateTime: Date,
) {
    companion object {
        fun of(configs: Map<String, String?>, articleCount: Int, categoryCount: Int, tagCount: Int, user: User, links: List<Link>, updateTime: Date): BlogVO {
            return BlogVO(
                configs[Config.KEY_SITE_DOMAIN] ?: "domain.org",
                configs[Config.KEY_SITE_TITLE] ?: "Blog",
                configs[Config.KEY_SITE_DESCRIPTION] ?: "",
                configs[Config.KEY_SITE_KEYWORDS] ?: "",
                configs[Config.KEY_SITE_BEIAN_ICP] ?: "",
                configs[Config.KEY_SITE_BEIAN_PS] ?: "",
                configs[Config.KEY_SITE_BAIDU] ?: "",
                configs[Config.KEY_SITE_CREATE_YEAR] ?: "2020",

                articleCount,
                categoryCount,
                tagCount,

                user.avatar ?: "",
                user.nickname ?: "",
                user.description ?: "",
                user.email ?: "",

                user.accounts?.github ?: "",
                user.accounts?.weibo ?: "",
                user.accounts?.google ?: "",
                user.accounts?.twitter ?: "",
                user.accounts?.facebook ?: "",
                user.accounts?.stackOverflow ?: "",
                user.accounts?.youtube ?: "",
                user.accounts?.instagram ?: "",
                user.accounts?.skype ?: "",

                links.map { LinkVO.of(it) },

                updateTime,
            )
        }
    }
}
