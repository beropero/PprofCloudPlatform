<template>
  <v-app
    class="v-data-table-container"
    style="background-color: white;color: black;"
  >
    <v-data-table-server
      v-model:items-per-page="itemsPerPage"
      
      :headers="headers as any" 
      :items="serverItems"
      :items-length="totalItems"
      :loading="loading"
      :search="search"
      item-value="name"
      style="background-color: white;color: black;"
      @update:options="loadItems"
    >
      <template #top>
        <v-breadcrumbs :items="['我的项目','服务模块','测试文件']" />
        <tr>
          <td>
            <v-text-field
              v-model="id"
              class="ma-2"
              density="compact"
              placeholder="Search Id..."
              hide-details
              width="200"
            />
          </td>
          <td>
            <v-text-field
              v-model="ptype"
              class="ma-2"
              density="compact"
              placeholder="Search Name..."
              hide-details
              width="200"
            />
          </td>
          <td>
            <v-text-field
              v-model="comment"
              class="ma-2"
              density="compact"
              placeholder="Search Description..."
              hide-details
              width="200"
            />
          </td>
          <v-divider />
        </tr>
      </template>
      <!-- info按钮 -->
      <template #[`item.action`]="item">
        <v-btn
          density="compact"
          icon="mdi-chevron-right"
          @click="() => router.push({ path: '/user/profile', query: { projectId: (item as any).item.mi } })"
        />
        <v-divider
          class="ms-3"
          inset
          vertical
        />
        <v-btn
          density="compact"
          icon="mdi-delete-outline"
        />
      </template>
    </v-data-table-server>
  </v-app>
</template>
<script lang="ts">
    import { GetProfileByPageUser } from '@/api/profile';
    import router from '@/router';
    export default {
      data: () => ({
        itemsPerPage: 10,
        headers: [
          {
            title: 'ID',
            align: 'start',
            sortable: false,
            key: 'microserviceId',
          },
          { title: 'Type', key: 'microserviceName', sortable: false, align: 'start' },
          { title: 'Comment', key: 'microserviceDescription', sortable: false, align: 'start' },
          { title: '', key: 'action', sortable: false, align: 'end' },
        ],
        serverItems: [],
        loading: true,
        totalItems: 0,
        comment: '',
        id: '',
        search: '',
        description: '',
        ptype: '',
        router,
        // 读取参数
        MicroserviceId: router.currentRoute.value.query.microserviceId,
      }),
      watch: {
        name () {
          this.search = String(Date.now())
        },
        id () {
          this.search = String(Date.now())
        },
        description () {
          this.search = String(Date.now())
        },
        ip () {
          this.search = String(Date.now())
        },
        port () {
          this.search = String(Date.now())
        },
      },
      methods: {
        async loadItems ({ page, itemsPerPage }: { page: number, itemsPerPage: number, sortBy: { key: string, order: string }[] }) {
            this.loading = true
  
            await GetProfileByPageUser({
                page: page,
                pageSize: itemsPerPage,
                microserviceId: this.MicroserviceId as unknown as number,
                comment: this.comment,
                ptype: this.ptype,
            }).then((res) => {
              const data = res.data.data || []
              console.log(data)
              this.serverItems = data.microservices || []
              this.totalItems = data.total || 0
              this.loading = false
            })
           
        },
      },
    }
</script>
  
<style scoped>
.v-data-table-container {
    overflow-y: auto;
}
</style>