package com.liux.blog.markdown.parser

import com.liux.blog.markdown.Flag
import com.liux.blog.markdown.bean.CatalogueItem
import com.liux.blog.markdown.node.*
import org.commonmark.internal.BlockContinueImpl
import org.commonmark.node.Block
import org.commonmark.node.CustomBlock
import org.commonmark.parser.block.*

class NexTNodeParser(
    private val block: CustomBlock
) : AbstractBlockParser() {

    companion object {
        private const val CUSTOM_NODE_TOC = "[toc]"
        private const val CUSTOM_NODE_MORE = "<!--more-->"
    }

    override fun getBlock(): Block {
        return block
    }

    override fun tryContinue(parserState: ParserState?): BlockContinue? {
        return when (block) {
            is MoreSuspend -> {
                BlockContinueImpl(-1, -1, false)
            }
            else -> BlockContinue.none()
        }
    }

    class Factory(
        private val flag: Flag,
        private val catalogues: List<CatalogueItem>,
    ) : AbstractBlockParserFactory() {
        override fun tryStart(state: ParserState, matchedBlockParser: MatchedBlockParser): BlockStart? {
            if (state.nextNonSpaceIndex > 0) return BlockStart.none()

            if (CUSTOM_NODE_TOC[0] == state.line.content[0] && CUSTOM_NODE_TOC == state.line.content.toString().lowercase()) {
                return if (flag.showToc()) {
                    BlockStart.of(NexTNodeParser(Catalogue(catalogues))).atIndex(state.index)
                } else {
                    BlockStart.of(NexTNodeParser(NoneCatalogue())).atIndex(state.index)
                }
            }

            if (CUSTOM_NODE_MORE[0] == state.line.content[0] && CUSTOM_NODE_MORE == state.line.content.toString().replace(" ", "").lowercase()) {
                if (flag.showMoreAnchor()) {
                    return BlockStart.of(NexTNodeParser(MoreAnchor())).atIndex(state.index)
                }
                if (flag.isMoreSuspend()) {
                    return BlockStart.of(NexTNodeParser(MoreSuspend())).atIndex(state.index)
                }
                return BlockStart.of(NexTNodeParser(NoneMore())).atIndex(state.index)
            }

            return BlockStart.none()
        }
    }
}
