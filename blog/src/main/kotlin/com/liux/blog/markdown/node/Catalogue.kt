package com.liux.blog.markdown.node

import com.liux.blog.markdown.bean.CatalogueItem
import org.commonmark.node.CustomBlock

class Catalogue(
    val catalogues: List<CatalogueItem>
) : CustomBlock()
