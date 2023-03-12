package org.commonmark.ext.next

import org.commonmark.internal.BlockContinueImpl
import org.commonmark.node.Block
import org.commonmark.node.CustomBlock
import org.commonmark.parser.block.*

class NexTBlockParser(
    private val block: CustomBlock
) : AbstractBlockParser() {

    override fun getBlock(): Block {
        return block
    }

    override fun tryContinue(parserState: ParserState?): BlockContinue? {
        return when (block) {
            is SuspendMoreBlock -> {
                BlockContinueImpl(-1, -1, false)
            }
            else -> BlockContinue.finished()
        }
    }

    class Factory(
        private val flag: Flag
    ) : AbstractBlockParserFactory() {
        override fun tryStart(state: ParserState, matchedBlockParser: MatchedBlockParser): BlockStart? {
            if (state.nextNonSpaceIndex > 0) return BlockStart.none()
            if (CUSTOM_NODE_TOC[0] == state.line.content[0] && CUSTOM_NODE_TOC == state.line.content.toString().lowercase()) {
                if (flag.isShowToc()) {
                    return BlockStart.of(NexTBlockParser(ShowTocBlock())).atIndex(state.index)
                }
                return BlockStart.of(NexTBlockParser(TocBlock())).atIndex(state.index)
            }
            if (CUSTOM_NODE_MORE[0] == state.line.content[0] && CUSTOM_NODE_MORE == state.line.content.toString().replace(" ", "").lowercase()) {
                if (flag.isMoreSuspend()) {
                    return BlockStart.of(NexTBlockParser(SuspendMoreBlock())).atIndex(state.index)
                }
                if (flag.isMoreAnchor()) {
                    return BlockStart.of(NexTBlockParser(AnchorMoreBlock())).atIndex(state.index)
                }
                return BlockStart.of(NexTBlockParser(MoreBlock())).atIndex(state.index)
            }
            return BlockStart.none()
        }
    }
}
