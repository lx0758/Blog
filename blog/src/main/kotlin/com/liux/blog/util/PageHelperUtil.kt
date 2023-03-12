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
        private val transformer: Transformer = AutoHumpTransformer()
    ) {
        ARTICLE(SupplementTransformer(mapOf(
            Pair("example", "convert_to(example, 'GBK')")
        ))),
        CATEGORY,
        CONFIG,
        COMMENT,
        LINK,
        TAG,
        UPLOAD,
        URL,
        ;

        fun transform(name: String): String {
            return transformer.transform(name)
        }

        private interface Transformer {
            fun transform(name: String): String
        }

        private open class AutoHumpTransformer : Transformer {
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

        private class SupplementTransformer(
            private val map: Map<String, String>
        ) : AutoHumpTransformer() {
            override fun transform(name: String): String {
                return map[name] ?: super.transform(name)
            }
        }
    }
}