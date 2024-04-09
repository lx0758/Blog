package com.liux.blog.markdown.parser

import com.liux.blog.markdown.Flag
import com.liux.blog.markdown.bean.CatalogueItem
import org.commonmark.ext.heading.anchor.IdGenerator

import org.commonmark.parser.Parser

class NexTParseExtension(
    private val flag: Flag,
    private val idGenerator: IdGenerator,
    private val catalogueResult : MutableList<CatalogueItem>?,
) : Parser.ParserExtension {
    override fun extend(parserBuilder: Parser.Builder) {
        val catalogueStore = NexTCataloguePostProcessor.CatalogueStore(idGenerator, catalogueResult)
        parserBuilder.customBlockParserFactory(NexTNodeParser.Factory(flag, catalogueStore.getCatalogueList()))
        parserBuilder.postProcessor(NexTCataloguePostProcessor(catalogueStore))
    }
}
