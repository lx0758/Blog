<template>
  <el-container class="full-page" direction="vertical">
    <el-space wrap>
      <el-input placeholder="请输入文章ID" v-model="state.filter.article" clearable />
      <el-input placeholder="请输入作者" v-model="state.filter.author" clearable />
      <el-input placeholder="请输入内容" v-model="state.filter.content" clearable />
      <el-input placeholder="请输入邮箱" v-model="state.filter.email" clearable />
      <el-input placeholder="请输入IP地址" v-model="state.filter.ip" clearable />
      <blog-selector v-model="state.filter.status" v-bind:type="5" />
      <el-button type="primary" plain icon="Search" @click="onFilterSearch">搜索</el-button>
      <el-button type="info" plain icon="Delete" @click="onFilterClear">清空</el-button>
    </el-space>

    <el-divider />

    <el-table class="full-page" stripe :data="state.data.list" @sort-change="onSortChange">
      <el-table-column prop="id" label="ID" width="70" fixed sortable="custom"></el-table-column>
      <el-table-column prop="articleName" label="文章" width="180" fixed>
        <template #default="scope">
          <el-link v-if="scope.row.articleUrl" :href="'/article/' + scope.row.articleUrl" type="primary" target="_blank">
            {{ scope.row.articleTitle }}
          </el-link>
          <el-link v-else :href="'/article/' + scope.row.articleId" type="primary" target="_blank">
            {{ scope.row.articleTitle }}
          </el-link>
        </template>
      </el-table-column>
      <el-table-column prop="nickname" label="昵称" width="100">
        <template #default="scope">
          <el-tooltip
              class="box-item"
              effect="dark"
              placement="top"
              :content="scope.row.nickname">
            {{ scope.row.nickname }}
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column prop="email" label="邮箱" width="150">
        <template #default="scope">
          <el-tooltip
              class="box-item"
              effect="dark"
              placement="top"
              :content="scope.row.email">
            {{ scope.row.email }}
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column prop="ip,location" label="IP" width="130">
        <template #default="scope">
          <el-tooltip class="box-item" effect="dark" :content="scope.row.location" placement="top">
            <el-link :href="'https://www.ip138.com/iplookup.php?ip=' + scope.row.ip" type="primary" target="_blank">
              {{ scope.row.ip }}
            </el-link>
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column prop="content" label="内容" min-width="300">
        <template #default="scope">
          <el-tooltip
              class="box-item"
              effect="dark"
              placement="top">
            <template #content>
              <pre>{{ scope.row.content }}</pre>
            </template>
            <p class="single-line">
              {{ scope.row.content }}
            </p>
          </el-tooltip>
        </template>
      </el-table-column>
      <date-time-column prop="createTime" label="创建时间" width="170" format="YYYY-MM-DD HH:mm:ss"></date-time-column>
      <date-time-column prop="updateTime" label="更新时间" width="170" format="YYYY-MM-DD HH:mm:ss"></date-time-column>
      <el-table-column prop="status" label="状态" width="90">
        <template #default="scope">
          <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'" disable-transitions size="small">
            {{ scope.row.status === 1 ? '已发布' : '待审核' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="160">
        <template #default="scope">
          <el-button plain size="small" @click="onShowReplyComment(scope.row)" v-if="scope.row.status === 1">回复</el-button>
          <el-button plain size="small" @click="onShowVerifyComment(scope.row)" v-if="scope.row.status === 0">审核</el-button>
          <el-button plain type="danger" size="small" @click="onDeleteComment(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination :page-count="state.data.pages" @current-change="onCurrentPageChange"
      style="text-align: center; margin-top: 20px" background layout="prev, pager, next" />
  </el-container>

  <el-dialog title="审核评论" v-model="verifyCommentDialogState.isShow" :close-on-click-modal="false">
    <el-descriptions direction="horizontal" :column="3" border>
      <el-descriptions-item label="文章" :span="2">{{
        verifyCommentDialogState.formModel.articleTitle
      }}
      </el-descriptions-item>
      <el-descriptions-item label="时间" :span="1">{{ verifyCommentDialogState.formModel.time }}</el-descriptions-item>
      <el-descriptions-item label="昵称" :span="1">{{
        verifyCommentDialogState.formModel.nickname
      }}
      </el-descriptions-item>
      <el-descriptions-item label="邮箱" :span="1">{{ verifyCommentDialogState.formModel.email }}</el-descriptions-item>
      <el-descriptions-item label="链接" :span="1">{{ verifyCommentDialogState.formModel.url }}</el-descriptions-item>
      <el-descriptions-item label="IP" :span="1">{{ verifyCommentDialogState.formModel.ip }}</el-descriptions-item>
      <el-descriptions-item label="浏览器" :span="1">{{
        verifyCommentDialogState.formModel.browser
      }}
      </el-descriptions-item>
      <el-descriptions-item label="系统" :span="1">{{
        verifyCommentDialogState.formModel.system
      }}
      </el-descriptions-item>
    </el-descriptions>
    <div style="
        margin: 20px 0 20px 0;
        word-break: normal;
        width: auto;
        white-space: pre-wrap;
        word-wrap: break-word;
        overflow: hidden;
      ">
      <span>{{ verifyCommentDialogState.formModel.content }}</span>
    </div>
    <el-button type="primary" @click="onVerifyComment">通过</el-button>
    <el-button @click="verifyCommentDialogState.hide()">取消</el-button>
  </el-dialog>

  <el-dialog title="回复评论" v-model="replayCommentDialogState.isShow" :close-on-click-modal="false">
    <el-descriptions direction="horizontal" :column="3" border>
      <el-descriptions-item label="文章" :span="2">{{
        replayCommentDialogState.formModel.articleTitle
      }}
      </el-descriptions-item>
      <el-descriptions-item label="时间" :span="1">{{ replayCommentDialogState.formModel.time }}</el-descriptions-item>
      <el-descriptions-item label="昵称" :span="1">{{
        replayCommentDialogState.formModel.nickname
      }}
      </el-descriptions-item>
      <el-descriptions-item label="邮箱" :span="1">{{ replayCommentDialogState.formModel.email }}</el-descriptions-item>
      <el-descriptions-item label="链接" :span="1">{{ replayCommentDialogState.formModel.url }}</el-descriptions-item>
      <el-descriptions-item label="IP" :span="1">{{
        replayCommentDialogState.formModel.ip + '/' + replayCommentDialogState.formModel.location
      }}
      </el-descriptions-item>
      <el-descriptions-item label="浏览器" :span="1">{{
        replayCommentDialogState.formModel.browser
      }}
      </el-descriptions-item>
      <el-descriptions-item label="系统" :span="1">{{
        replayCommentDialogState.formModel.system
      }}
      </el-descriptions-item>
    </el-descriptions>
    <div style="
        margin: 20px 0 20px 0;
        word-break: normal;
        width: auto;
        white-space: pre-wrap;
        word-wrap: break-word;
        overflow: hidden;
      ">
      <span>{{ replayCommentDialogState.formModel.content }}</span>
    </div>
    <el-input type="textarea" rows="6" v-model="replayCommentDialogState.formModel.replayContent"
      :placeholder="'@' + replayCommentDialogState.formModel.nickname" />
    <div style="height: 20px" />
    <el-button type="primary" @click="onReplyComment(true)">回复并邮件通知</el-button>
    <el-button type="primary" @click="onReplyComment(false)">仅回复不通知</el-button>
    <el-button @click="replayCommentDialogState.hide()">取消</el-button>
  </el-dialog>
</template>

<script setup lang="ts">
import dayjs from 'dayjs'
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ListState, DialogState } from '@/bean'
import { addCommentByReplay, queryComment, updateCommentToVerify, deleteComment } from '@/api'

interface CommentFilter {
  article: string | null
  author: string | null
  content: string | null
  email: string | null
  ip: string | null
  status: number | null
}

interface OptionCommentState {
  parentId: number
  articleId: number
  articleTitle: string
  articleUrl: string
  time: string
  nickname: string
  email: string
  url: string
  ip: string
  location: string
  browser: string
  system: string
  content: string

  replayContent: string
}

const state = ref(new ListState<CommentFilter>({} as CommentFilter))

const verifyCommentDialogStateFormRef = ref()
const verifyCommentDialogState = ref(new DialogState<OptionCommentState>(
  verifyCommentDialogStateFormRef,
  {} as OptionCommentState,
  {},
))

const replayCommentDialogStateFormRef = ref()
const replayCommentDialogState = ref(new DialogState<OptionCommentState>(
  replayCommentDialogStateFormRef,
  {} as OptionCommentState,
  {},
))

function onRefresh() {
  queryComment(
    state.value.filter.article,
    state.value.filter.author,
    state.value.filter.content,
    state.value.filter.email,
    state.value.filter.ip,
    state.value.filter.status,
    state.value.data.pageNum,
    state.value.data.pageSize,
    state.value.order.name,
    state.value.order.method
  ).then((data) => {
    state.value.data = data.data
  })
}

function onFilterSearch() {
  onRefresh()
}

function onFilterClear() {
  state.value.clearFilter()
  onRefresh()
}

function onSortChange(column: any) {
  state.value.saveOrder(column.prop, column.order)
  onRefresh()
}

function onCurrentPageChange(currentPage: number) {
  const changed = state.value.data.pageNum != currentPage
  state.value.data.pageNum = currentPage
  if (changed) {
    onRefresh()
  }
}

function onShowReplyComment(row: any) {
  replayCommentDialogState.value.formModel = {
    articleId: row.articleId,
    parentId: row.id,
    articleTitle: row.articleTitle,
    articleUrl: row.articleUrl,
    nickname: row.nickname,
    email: row.email,
    url: row.url,
    ip: row.ip,
    location: row.location,
    browser: row.browser,
    system: row.system,
    time: dayjs(row.createTime).format('YYYY-MM-DD HH:mm:ss'),
    content: row.content,
    replayContent: ''
  }
  replayCommentDialogState.value.show()
}

function onReplyComment(emailNotify: boolean) {
  const replayData = replayCommentDialogState.value.formModel
  if (!replayData.replayContent || replayData.replayContent === '') {
    ElMessage.warning('回复内容不能为空')
    return
  }
  addCommentByReplay(
    replayData.articleId,
    replayData.parentId,
    '@' + replayData.nickname + '\n' + replayData.replayContent,
    emailNotify
  ).then(() => {
    ElMessage.success('回复成功')
    replayCommentDialogState.value.hide()
    onRefresh()
  })
}

function onShowVerifyComment(row: any) {
  verifyCommentDialogState.value.formModel = {
    articleId: row.id,
    articleTitle: row.articleTitle,
    nickname: row.nickname,
    email: row.email,
    url: row.url,
    ip: row.ip,
    location: row.location,
    browser: row.browser,
    system: row.system,
    time: dayjs(row.createTime).format('YYYY-MM-DD HH:mm:ss'),
    content: row.content
  } as OptionCommentState
  verifyCommentDialogState.value.show()
}

function onVerifyComment() {
  const verifyData = verifyCommentDialogState.value.formModel
  updateCommentToVerify(verifyData.articleId).then(() => {
    ElMessage.success('审核成功')
    verifyCommentDialogState.value.hide()
    onRefresh()
  })
}

function onDeleteComment(row: any) {
  ElMessageBox.confirm('确认删除?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteComment(row.id).then(() => {
      ElMessage.success('删除成功')
      onRefresh()
    })
  })
}

onMounted(() => {
  onRefresh()
})
</script>

<style>
.single-line {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>
