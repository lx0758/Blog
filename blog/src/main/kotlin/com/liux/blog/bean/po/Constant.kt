package com.liux.blog.bean.po

// 删除状态
const val STATE_DELETED = -1
// 文章/文件状态
const val STATE_UNPUBLISHED = 0
const val STATE_PUBLISHED = 1
// 评论状态
const val STATE_PENDING = 0
const val STATE_AUDITED = 1
// 用户状态
const val STATE_NOT_ACTIVATED = 0
const val STATE_ACTIVATED = 1

// 文章格式
const val FORMAT_MARKDOWN = 0
const val FORMAT_HTML = 1

// 配置键值
const val SITE_DOMAIN = "SITE_DOMAIN"
const val SITE_TITLE = "SITE_TITLE"
const val SITE_DESCRIPTION = "SITE_DESCRIPTION"
const val SITE_KEYWORDS = "SITE_KEYWORDS"
const val SITE_BEIAN = "SITE_BEIAN"
const val SITE_BAIDU = "SITE_BAIDU"
const val SITE_CREATE_YEAR = "SITE_CREATE_YEAR"
