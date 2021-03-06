package com.liux.blog.markdown

import com.liux.blog.bean.vo.CatalogueVO
import org.commonmark.node.Document

class NexTDocument(
    val original: Document,
    val catalogues: Collection<CatalogueVO>
) : Document()