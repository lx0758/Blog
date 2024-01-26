<template>
  <el-table-column :formatter="onFormatDate" :prop="props.prop" :label="props.label" :width="props.width"
    sortable="custom"></el-table-column>
</template>

<script setup lang="ts">
import dayjs from 'dayjs'

const props = defineProps(['prop', 'label', 'width', 'format'])

function onFormatDate(row: any, column: any) {
  const date = row[column.property]
  if (date == null) return '-'
  return dayjs(date).format(props.format)
}
</script>

<style></style>
