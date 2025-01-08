<template>
  <el-container :style="'height: ' + height + 'px;'">
    <slot />
  </el-container>
</template>

<script setup lang="ts">
import { ref, nextTick, onMounted, onUnmounted } from 'vue'

const props = defineProps(['bottom', 'min'])

const height = ref(calculateNewHeight())

function calculateNewHeight(): number {
  let newHeight = window.innerHeight - props.bottom
  if (newHeight < props.min) {
    newHeight = props.min
  }
  return newHeight
}

onMounted(() => {
  nextTick(() => {
    window.onresize = () => {
      height.value = calculateNewHeight()
    }
  })
})
onUnmounted(() => {
  nextTick(() => {
    window.onresize = null
  })
})
</script>

<style></style>
