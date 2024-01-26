import { ref, type Ref } from 'vue'
import type { FormInstance, FormRules } from 'element-plus'

interface ListOrder {
    name: string | null
    method: string | null
}
interface ListData {
    pageNum: number
    pageSize: number
    pages: number
    total: number
    list: Array<any>
}
export class ListState<T extends Object> {
    filter: T

    order: ListOrder = {
        name: null,
        method: null,
    }

    data = ref<ListData>({
        pageNum: 1,
        pageSize: 20,
        pages: 0,
        total: 0,
        list: [],
    })

    constructor(filter: T) {
        this.filter = filter
    }

    clearFilter() {
        this.filter = {} as T
    }

    saveOrder(name: string | null, method: string | null) {
        this.order.name = name
        this.order.method = method
    }
}

export class DialogState<T extends Object> {
    isShow: boolean = false
    formRef: Ref<FormInstance | undefined>
    formModel: T
    formRules: FormRules<T>

    constructor(formRef: Ref<FormInstance | undefined>, form: T, rules: FormRules<T>) {
        this.formRef = formRef
        this.formModel = form
        this.formRules = rules
    }

    show() {
        this.isShow = true
    }
    reset() {
        if (this.formRef != undefined && this.formRef.value != undefined) {
            this.formRef.value.resetFields()
        }
        this.formModel = {} as T
    }
    hide() {
        this.isShow = false
    }
}