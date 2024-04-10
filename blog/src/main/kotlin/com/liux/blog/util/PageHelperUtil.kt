package com.liux.blog.util

import com.github.pagehelper.PageHelper

class PageHelperUtil {

    companion object {
        fun orderBy(type: Type, orderName: String?, orderMethod: String?) {
            if (orderName == null) return
            val name = type.transform(orderName)
            val method = if ("descending" != orderMethod)
                "ASC NULLS FIRST"
            else
                "DESC NULLS LAST"
            PageHelper.orderBy("$name $method")
        }
    }

    enum class Type(
        vararg pairs: Pair<String, String>
    ) {
        ARTICLE(
            // 按需添加自定义规则
            "column_name" to "convert_to(field_name, 'GBK')",
        ),
        CATEGORY,
        CONFIG,
        COMMENT,
        LINK,
        TAG,
        UPLOAD,
        URL,
        ;

        private val transformer = CustomRuleTransformer(*pairs)

        fun transform(name: String): String {
            return transformer.transform(name)
        }

        private interface Transformer {
            fun transform(name: String): String
        }

        private open class AutoConvertTransformer : Transformer {
            override fun transform(name: String): String {
                val builder = StringBuilder()
                name.forEach {
                    if (Character.isUpperCase(it)) {
                        builder.append("_")
                        builder.append(it.lowercase())
                    } else {
                        builder.append(it)
                    }
                }
                return builder.toString()
            }
        }

        private open class CustomRuleTransformer(
            vararg pairs: Pair<String, String>
        ) : AutoConvertTransformer() {
            private val map = mapOf(*pairs)
            override fun transform(name: String): String {
                return map[name] ?: super.transform(name)
            }
        }
    }
}