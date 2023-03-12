<template>
  <el-select :model-value="value" @change="onSelected" :placeholder="name" :size="size">
    <el-option
        v-for="item in options"
        :key="item.value"
        :label="item.label"
        :value="item.value">
    </el-option>
  </el-select>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import {queryCategoryOptions} from "@/api";

export default defineComponent({
  props: {
    value: Number,
    type: Number,
  },
  mounted() {
    switch (this.type) {
      case 1:
        this.name = "文章分类"
        queryCategoryOptions()
            .then(data => {
              this.options = data
            })
        break
      case 2:
        this.name = "文章格式"
        this.options = [
          {
            value: 0,
            label: 'Markdown',
          },
          {
            value: 1,
            label: 'Html',
          },
        ]
        break
      case 3:
        this.name = "允许评论"
        this.options = [
          {
            value: 0,
            label: '禁止',
          },
          {
            value: 1,
            label: '允许',
          },
        ]
        break
      case 4:
        this.name = "文章状态"
        this.options = [
          {
            value: 0,
            label: '草稿',
          },
          {
            value: 1,
            label: '发布',
          },
        ]
        break
      case 5:
        this.name = "评论状态"
        this.options = [
          {
            value: 0,
            label: '待审核',
          },
          {
            value: 1,
            label: '已审核',
          },
        ]
        break
      case 6:
        this.name = "友链状态"
        this.options = [
          {
            value: 0,
            label: '禁用',
          },
          {
            value: 1,
            label: '启用',
          },
        ]
        break
      case 7:
        this.name = "短链状态"
        this.options = [
          {
            value: 0,
            label: '禁用',
          },
          {
            value: 1,
            label: '启用',
          },
        ]
        break
    }
  },
  data() {
    return {
      name: '',
      size: this.$attrs.size,
      options: {},
    }
  },
  methods: {
    onSelected(value: any) {
      this.$emit('update:value', value)
    }
  }
})
</script>

<style lang="less">
</style>