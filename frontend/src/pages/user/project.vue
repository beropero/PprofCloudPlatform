<template>
  <v-app
    class="v-data-table-container"
    style="background-color: white;color: black;"
  >
    <!--创建对话框-->
    <v-dialog
      v-model="createshow"
      max-width="500"
      persistent
    >
      <v-card
        style="background-color: #fff;color: #111;"
        title="新建项目"
      > 
        <v-divider />
        <v-card-text class="px-5">
          <v-form
            fast-fail
          >
            <v-text-field
              v-model="creatforms.name"
              label="name"
            />

            <v-text-field
              v-model="creatforms.description"
              label="description"
            />
          </v-form>
        </v-card-text>
  
        <!--按钮-->
        <v-card-actions>
          <v-spacer />

          <v-btn
            color="red"
            text="Close"
            variant="plain"
            @click="createshow = false"
          />

          <v-btn
          
            text="Save"
            variant="plain"
            @click="createproject()"
          />
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!--数据表格-->
    <v-data-table-server
      v-model:items-per-page="itemsPerPage"
      :headers="headers as any" 
      :items="serverItems"
      :items-length="totalItems"
      :loading="loading"
      :search="search"
      item-value="name"
      style="background-color: white;color: black;"
      :hide-no-data="true"
      @update:options="loadItems"
    >
      <template #top>
        <v-breadcrumbs :items="['我的项目','']" />
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
              v-model="name"
              class="ma-2"
              density="compact"
              placeholder="Search Name..."
              hide-details
              width="200"
            />
          </td>
          <td>
            <v-text-field
              v-model="description"
              class="ma-2"
              density="compact"
              placeholder="Search Description..."
              hide-details
              width="200"
            />
          </td>
          <td>
            <v-btn
              v-model="description"
              density="compact"
              icon="mdi-plus"
              style="bottom: 15px; left: 10px;"
              @click="createshow = true"
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
          @click="() => router.push({ path: '/user/microservice', query: { projectId: (item as any).item.projectId } })"
        />
        <v-divider
          class="ms-3"
          inset
          vertical
        />
        <v-speed-dial
          location="left center"
          transition="fade-transition"
        >
          <template #activator="{ props: activatorProps }">
            <v-fab
              v-bind="activatorProps"
              density="compact"
              icon="mdi-dots-vertical"
            />
          </template>

          <v-btn
            key="1"
            icon="mdi-delete"
          />
          <v-btn
            key="2"
            icon="mdi-pencil"
          />
        </v-speed-dial>
      </template>
    </v-data-table-server>
  </v-app>
</template>
<script lang="ts">
  import { CreateProject, GetProjectByPageUser } from '@/api/project';
  import router from '@/router';
  export default {
    data: () => ({
      itemsPerPage: 10,
      headers: [
        {
          title: 'ID',
          align: 'start',
          sortable: false,
          key: 'projectId',
        },
        { title: 'Name', key: 'projectName', sortable: false, align: 'start' },
        { title: 'Description', key: 'projectDescription', sortable: false, align: 'start' },
        { title: 'Token', key: 'projectToken', sortable: false, align: 'start' },
        { title: '', key: 'action', sortable: false, align: 'end' },
      ],
      serverItems: [],
      loading: true,
      totalItems: 0,
      id: '',
      name: '',
      search: '',
      description: '',
      router,
      createshow: false,
      creatforms: {
        name: '',
        description: '',
      },
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
    },
    methods: {
      async loadItems ({ page, itemsPerPage }: { page: number, itemsPerPage: number, sortBy: { key: string, order: string }[] }) {
          this.loading = true

          await GetProjectByPageUser({
            page: page,
            pageSize: itemsPerPage,
            projectDescription: this.description,
            projectId: this.id as unknown as number,
            projectName: this.name,
          }).then((res) => {
            const data = res.data.data || []
            this.serverItems = data.projects || []
            this.totalItems = data.total || 0
            this.loading = false
          })
         
      },
      async createproject () {
          await CreateProject({
            projectName: this.creatforms.name,
            projectDescription: this.creatforms.description,
          }).then((res) => {
            const data = res.data.data || []
            this.id = data.id
          })
         this.createshow = false
         
      }
    },

  }
</script>

<style scoped>
.v-data-table-container {
    overflow-y: auto;
}
</style>