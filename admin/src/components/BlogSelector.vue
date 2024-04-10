<template>
  <el-select
    v-model="value"
    style="width: 180px"
    :placeholder="stateRef.placeholder"
    :size="size">
    <el-option v-for="item in stateRef.options" :key="item.value" :label="item.label" :value="item.value">
    </el-option>
  </el-select>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { queryCategoryOptions } from '@/api'

interface Option {
    label: string
    value: any
}
interface State {
  placeholder: string
  options: Array<Option>
}

const props = defineProps(['value', 'type', 'size'])
const emits = defineEmits(['update:value'])

const value = computed({
  get() {
    return props.value
  },
  set(value) {
    emits('update:value', value)
  }
})

const stateRef = ref<State>({
  placeholder: '',
  options: []
})

onMounted(() => {
  switch (props.type) {
    case 1:
      stateRef.value.placeholder = '文章分类'
      queryCategoryOptions().then((data) => {
        stateRef.value.options = data
      })
      break
    case 2:
      stateRef.value.placeholder = '文章格式'
      stateRef.value.options = [
        {value: 0, label: 'Markdown'},
        {value: 1, label: 'Html'}
      ]
      break
    case 3:
      stateRef.value.placeholder = '允许评论'
      stateRef.value.options = [
        {value: false, label: '禁止'},
        {value: true, label: '允许'}
      ]
      break
    case 4:
      stateRef.value.placeholder = '文章状态'
      stateRef.value.options = [
        {value: 0, label: '草稿'},
        {value: 1, label: '发布'}
      ]
      break
    case 5:
      stateRef.value.placeholder = '评论状态'
      stateRef.value.options = [
        {value: 0, label: '待审核'},
        {value: 1, label: '已审核'}
      ]
      break
    case 6:
      stateRef.value.placeholder = '友链状态'
      stateRef.value.options = [
        {value: 0, label: '禁用'},
        {value: 1, label: '启用'}
      ]
      break
    case 7:
      stateRef.value.placeholder = '短链状态'
      stateRef.value.options = [
        {value: 0, label: '禁用'},
        {value: 1, label: '启用'}
      ]
      break
    case 8:
      stateRef.value.placeholder = '片段状态'
      stateRef.value.options = [
        {value: 0, label: '禁用'},
        {value: 1, label: '启用'}
      ]
      break
  }
})
</script>

<style></style>
