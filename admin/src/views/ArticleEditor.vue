<template>

  <el-form label-width="80px">
    <el-row gutter="middle">
      <el-col :span="12">
        <el-form-item label="文章标题">
          <el-input
              v-model="title"
              placeholder="请输入文章标题"/>
        </el-form-item>
      </el-col>
      <el-col :span="12">
        <el-form-item label="文章分类">
          <blog-select
              v-model:value="categoryId"
              v-bind:type="1"
              style="width: 100%"/>
        </el-form-item>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12">
        <el-form-item label="文章链接">
          <el-input
              v-model="url"
              placeholder="请输入文章链接"/>
        </el-form-item>
      </el-col>
      <el-col :span="12">
        <el-form-item label="文章标签">
          <el-select
              v-model="tags"
              clearable
              multiple
              filterable
              allow-create
              default-first-option
              placeholder="请输入标签(回车添加)"
              style="width: 100%">
            <el-option v-for="item in tags" :key="item.value" :label="item.label" :value="item.label"></el-option>
          </el-select>
        </el-form-item>
      </el-col>
    </el-row>

    <el-container :style="'margin-bottom: 20px; height: ' + mainContentHeight + 'px'">
      <v-md-editor v-model="content"/>
    </el-container>

    <el-row>
      <el-col :span="12">
        <el-container>
          <el-form-item label="文章状态" style="width: 120px;">
            <el-switch
                v-model="status"
                :active-value="1"
                :inactive-value="0"/>
          </el-form-item>
          <el-form-item label="评论状态" style="width: 120px;">
            <el-switch
                v-model="enableComment"
                :active-value="true"
                :inactive-value="false"/>
          </el-form-item>
        </el-container>
      </el-col>
      <el-col :span="12">
        <el-container style="float: right;">
          <el-button type="info" plain icon="el-icon-back" @click="onBack">返回</el-button>
          <el-button type="success" plain icon="el-icon-document-remove" @click="onSaveArticle">保存</el-button>
          <el-button type="primary" plain icon="el-icon-document-checked" @click="onPublishArticle">发布</el-button>
        </el-container>
      </el-col>
    </el-row>

  </el-form>

</template>

<script lang="ts">
import {defineComponent} from 'vue';
import {queryArticleInfo, addArticle, updateArticle} from "@/api";
import BlogSelect from "@/components/BlogSelect.vue";

export default defineComponent({
  name: 'ArticleEditor',
  components: {
    BlogSelect,
  },
  mounted() {
    const params = this.$route.params
    if (params && params.articleId) {
      this.id = Number.parseInt(params.articleId as string)
      queryArticleInfo(this.id)
          .then(data => {
            data = data.data
            this.title = data.title
            this.categoryId = data.categoryId
            this.url = data.url
            this.tags = data.tags
            this.format = data.format
            this.content = data.content
            this.enableComment = data.enableComment
            this.status = data.status
          })
    }
    window.onresize = () => {
      this.onRefreshMainContentHeight()
    }
    this.onRefreshMainContentHeight()
  },
  unmounted() {
    window.onresize = null;
  },
  data() {
    const mainContentBottom = 300
    return {
      mainContentBottom,
      mainContentHeight: window.innerHeight - mainContentBottom,

      id: 0,
      title: '',
      categoryId: 0,
      url: null,
      tags: [],

      format: 0,
      content: '',

      weight: 0,
      enableComment: true,
      status: 0,
    }
  },
  methods: {
    onRefreshMainContentHeight() {
      this.mainContentHeight = window.innerHeight - this.mainContentBottom
    },

    onBack() {
      this.$router.back()
    },
    onSaveArticle() {
      if (this.id == 0) {
        addArticle(
            this.title,
            this.categoryId,
            this.url,
            this.tags,
            this.format,
            this.content,
            this.weight,
            this.enableComment,
            0,
        )
            .then(() => {
              this.$message.success("保存成功");
              this.onBack()
            })
      } else {
        updateArticle(
            this.id,
            this.title,
            this.categoryId,
            this.url,
            this.tags,
            this.format,
            this.content,
            this.weight,
            this.enableComment,
            this.status,
        )
            .then(() => {
              this.$message.success("更新成功");
              this.onBack()
            })
      }
    },
    onPublishArticle() {
      if (this.id == 0) {
        addArticle(
            this.title,
            this.categoryId,
            this.url,
            this.tags,
            this.format,
            this.content,
            this.weight,
            this.enableComment,
            1,
        )
            .then(() => {
              this.$message.success("发布成功");
              this.onBack()
            })
      } else {
        updateArticle(
            this.id,
            this.title,
            this.categoryId,
            this.url,
            this.tags,
            this.format,
            this.content,
            this.weight,
            this.enableComment,
            1,
        )
            .then(() => {
              this.$message.success("更新成功");
              this.onBack()
            })
      }
    },
  }
});
</script>

<style scoped lang="less">
</style>