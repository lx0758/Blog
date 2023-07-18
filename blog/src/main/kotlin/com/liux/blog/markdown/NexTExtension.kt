package com.liux.blog.markdown

import org.commonmark.Extension
import org.commonmark.parser.Parser
import org.commonmark.renderer.html.HtmlRenderer

class NexTExtension(private val flag: Flag) : Parser.ParserExtension, HtmlRenderer.HtmlRendererExtension {

    companion object {
        fun create(flag: Int = 0): Extension {
            return NexTExtension(Flag(flag))
        }
    }

    override fun extend(parserBuilder: Parser.Builder) {
        parserBuilder.customBlockParserFactory(NexTBlockParser.Factory(flag))
        parserBuilder.postProcessor(NexTPostProcessor())
    }

    override fun extend(htmlRendererBuilder: HtmlRenderer.Builder) {
        htmlRendererBuilder.nodeRendererFactory(NexTNodeRenderer.Factory())
        htmlRendererBuilder.attributeProviderFactory(NexTAttributeProvider.Factory())
    }
}
