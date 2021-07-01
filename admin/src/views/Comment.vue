<template>
  <el-container direction="vertical">

    <el-space wrap>
      <el-input
          placeholder="请输入文章ID"
          v-model="filter.article"
          clearable/>
      <el-input
          placeholder="请输入作者"
          v-model="filter.author"
          clearable/>
      <el-input
          placeholder="请输入内容"
          v-model="filter.content"
          clearable/>
      <el-input
          placeholder="请输入邮箱"
          v-model="filter.email"
          clearable/>
      <el-input
          placeholder="请输入IP地址"
          v-model="filter.ip"
          clearable/>
      <el-select v-model="filter.status" placeholder="状态">
        <el-option
            v-for="item in filter.statusOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value">
        </el-option>
      </el-select>
      <el-button type="primary" plain icon="el-icon-search" @click="onFilterSearch">搜索</el-button>
      <el-button type="info" plain icon="el-icon-delete" @click="onFilterClear">清空</el-button>
    </el-space>

    <el-divider/>

    <el-table
        :data="data.list"
        border
        stripe
        style="width: 100%; height: auto">
      <el-table-column prop="id" label="ID" width="60" fixed></el-table-column>
      <el-table-column prop="articleName" label="文章标题" min-width="200" :show-overflow-tooltip="true" fixed>
        <template #default="scope">
          <el-link :href="'/article/' + scope.row.articleId" type="primary" target="_blank">{{ scope.row.articleTitle }}</el-link>
        </template>
      </el-table-column>
      <el-table-column prop="nickname" label="昵称" width="120"></el-table-column>
      <el-table-column prop="email" label="邮箱" width="180"></el-table-column>
      <el-table-column prop="ip" label="IP" width="140"></el-table-column>
      <el-table-column prop="content" label="内容" min-width="300" :show-overflow-tooltip="true"></el-table-column>
      <el-table-column :formatter="onFormatDate" prop="createTime" label="创建时间" width="160"></el-table-column>
      <el-table-column :formatter="onFormatDate" prop="updateTime" label="更新时间" width="160"></el-table-column>
      <el-table-column prop="status" label="状态" width="70">
        <template #default="scope">
          <el-tag
              :type="scope.row.status === 1 ? 'success' : scope.row.status === -1 ? 'info' : 'danger'"
              disable-transitions
              size="medium">{{ scope.row.status === 1 ? '发布' : scope.row.status === -1 ? '删除' : '待审核'}}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="160">
        <template #default="scope">
          <el-button plain size="mini" @click="onShowReplyComment(scope.row)">回复</el-button>
          <el-button plain type="danger" size="mini" @click="onEditComment(scope.row)" v-if="scope.row.status === 0">审核</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
        :page-count="data.pages"
        @current-change="onCurrentPageChange"
        style="text-align: center; margin-top: 20px"
        background
        layout="prev, pager, next"/>

  </el-container>

  <el-dialog title="审核评论" v-model="verify">
    <el-form ref="verify" :model="verifyData" label-width="120px">
      <el-form-item label="文章标题" prop="title">
        <span>{{verifyData.articleTitle}}</span>
      </el-form-item>
      <el-form-item label="用户信息" prop="nickname">
        <span>{{verifyData.nickname + ' / ' + verifyData.email + ' / ' + verifyData.url}}</span>
      </el-form-item>
      <el-form-item label="设备信息" prop="nickname">
        <span>{{verifyData.ip + ' / ' + verifyData.browser + ' / ' + verifyData.system}}</span>
      </el-form-item>
      <el-form-item label="评论时间" prop="nickname">
        <span>{{verifyData.time}}</span>
      </el-form-item>
      <el-form-item label="评论内容" prop="nickname">
        <span>{{verifyData.content}}</span>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onVerifyComment">通过</el-button>
        <el-button type="primary" @click="onDeleteComment">删除</el-button>
        <el-button @click="verify = false">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>

  <el-dialog title="回复评论" v-model="replay">
    <el-form ref="replay" :model="replayData" label-width="120px">
      <el-form-item label="文章标题" prop="title">
        <span>{{replayData.articleTitle}}</span>
      </el-form-item>
      <el-form-item label="用户信息" prop="nickname">
        <span>{{replayData.nickname + ' / ' + replayData.email + ' / ' + replayData.url}}</span>
      </el-form-item>
      <el-form-item label="设备信息" prop="nickname">
        <span>{{replayData.ip + ' / ' + replayData.browser + ' / ' + replayData.system}}</span>
      </el-form-item>
      <el-form-item label="评论时间" prop="nickname">
        <span>{{replayData.time}}</span>
      </el-form-item>
      <el-form-item label="评论内容" prop="nickname">
        <span>{{replayData.content}}</span>
      </el-form-item>
      <el-form-item label="答复内容" prop="replayContent">
        <el-input type="textarea" v-model="replayData.replayContent"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onReplyComment">回复</el-button>
        <el-button @click="replay = false">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>

</template>

<script lang="ts">
import { defineComponent } from 'vue';
import dayjs from "dayjs"
import {addCommentByReplay, queryComment, updateCommentToVerify, deleteComment} from "@/api"

export default defineComponent({
  name: 'Comment',
  components: {},
  mounted() {
    this.onRefresh()
  },
  data() {
    return {
      filter: {
        article: null,
        author: null,
        content: null,
        email: null,
        ip: null,
        status: null,
        statusOptions: [
          {
            value: -1,
            label: '删除',
          },
          {
            value: 0,
            label: '草稿',
          },
          {
            value: 1,
            label: '发布',
          },
        ],
      },

      data: {
        pageNum: 1,
        pageSize: 20,
        pages: 0,
        total: 0,
        list: [],
      },

      verify: false,
      verifyData: {},

      replay: false,
      replayData: {},
    }
  },
  methods: {
    onFilterSearch() {
      this.onRefresh()
    },
    onFilterClear() {
      this.filter.article = null
      this.filter.author = null
      this.filter.content = null
      this.filter.email = null
      this.filter.ip = null
      this.filter.status = null
      this.onRefresh()
    },
    onFormatDate(row: any, column: any) {
      const date = row[column.property];
      return dayjs(date).format("YYYY-MM-DD HH:mm:ss");
    },
    onCurrentPageChange(currentPage: number) {
      this.data.pageNum = currentPage;
      this.onRefresh();
    },
    onShowReplyComment(row: any) {
      this.replayData = {
        articleId: row.articleId,
        parentId: row.id,

        articleTitle: row.articleTitle,
        nickname: row.nickname,
        email: row.email,
        url: row.url,
        ip: row.ip,
        browser: row.browser,
        system: row.system,
        time: dayjs(row.createTime).format("YYYY-MM-DD HH:mm:ss"),
        content: row.content,
      }
      this.replay = true
    },
    onReplyComment() {
      const replayData = this.replayData as any
      addCommentByReplay(
          replayData.articleId,
          replayData.parentId,
          replayData.replayContent,
      )
          .then(() => {
            this.$message.success("删除成功");
            this.onRefresh()
          })
    },
    onEditComment(row: any) {
      this.verifyData = {
        articleTitle: row.articleTitle,
        nickname: row.nickname,
        email: row.email,
        url: row.url,
        ip: row.ip,
        browser: row.browser,
        system: row.system,
        time: dayjs(row.createTime).format("YYYY-MM-DD HH:mm:ss"),
        content: row.content,
      }
      this.verify = true
    },
    onVerifyComment() {
      updateCommentToVerify((this.verifyData as any).id)
          .then(() => {
            this.$message.success("删除成功");
            this.onRefresh()
          })
    },
    onDeleteComment(row: any) {
      this.$confirm('确认删除?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteComment(row.id)
            .then(() => {
              this.$message.success("删除成功");
              this.onRefresh()
            })
      });
    },
    onRefresh() {
      queryComment(
          this.filter.article,
          this.filter.author,
          this.filter.content,
          this.filter.email,
          this.filter.ip,
          this.filter.status,
          this.data.pageNum,
          this.data.pageSize,
      )
          .then(data => {
            this.data = data.data;
          })
    },
  }
});
</script>

<style lang="less">
</style>