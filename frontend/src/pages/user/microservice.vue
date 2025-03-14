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
        title="新建服务"
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
            <v-row>
              <v-col cols="8">
                <v-text-field
                  v-model="creatforms.ip"
                  label="ip"
                />
              </v-col>
              <v-col cols="4">
                <v-text-field
                  v-model="creatforms.port"
                  label="port"
                  type="number"
                />
              </v-col>
            </v-row>
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
            @click="createmicroservice()"
          />
        </v-card-actions>
      </v-card>
    </v-dialog>
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
        <v-breadcrumbs :items="['我的项目','服务模块']" />
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
            <v-text-field
              v-model="ip"
              class="ma-2"
              density="compact"
              placeholder="Search Ip..."
              hide-details
              width="200"
            />
          </td>
          <td>
            <v-text-field
              v-model="port"
              class="ma-2"
              density="compact"
              placeholder="Search Port..."
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
          @click="() => router.push({ path: '/user/profile', query: { microserviceId: (item as any).item.microserviceId } })"
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
    import { CreateMicroservice, GetMicroserviceByPageUser } from '@/api/microservice';
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
          { title: 'Name', key: 'microserviceName', sortable: false, align: 'start' },
          { title: 'Description', key: 'microserviceDescription', sortable: false, align: 'start' },
          { title: 'Ip', key: 'ip', sortable: false, align: 'start' },
          { title: 'Port', key: 'port', sortable: false, align: 'start' },
          { title: 'Token', key: 'microserviceToken', sortable: false, align: 'start' },
          { title: '', key: 'action', sortable: false, align: 'end' },
        ],
        serverItems: [],
        loading: true,
        totalItems: 0,
        port: '',
        ip: '',
        id: '',
        name: '',
        search: '',
        description: '',
        router,
        createshow: false,
        creatforms:{
          name: '',
          description: '',
          ip:'',
          port:0,
        }
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
            await GetMicroserviceByPageUser({
                page: page,
                pageSize: itemsPerPage,
                Ip: this.ip,
                MicroserviceDescription: this.description,
                MicroserviceId: this.id as unknown as number,
                MicroserviceName: this.name,
                Port: this.port as unknown as number,
                ProjectId: router.currentRoute.value.query.projectId as unknown as number
            }).then((res) => {
              const data = res.data.data || []
              this.serverItems = data.microservices || []
              this.totalItems = data.total || 0
              this.loading = false
            })
        },
        async createmicroservice () {
            const ProjectId = router.currentRoute.value.query.projectId
            await CreateMicroservice({
                Ip: this.creatforms.ip,
                MicroserviceDescription: this.creatforms.description,
                MicroserviceName: this.creatforms.name,
                Port: this.creatforms.port,
                ProjectId: ProjectId as unknown as number
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