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
        title="新建测试"
      > 
        <v-divider />
        <v-card-text class="px-5">
          <v-form
            fast-fail
          >
            <v-text-field
              v-model="creatforms.microserviceName"
              label="Microservice Name"
              :disabled="true"
            />
            <v-row>
              <v-col cols="8">
                <v-text-field
                  v-model="creatforms.ip"
                  label="ip"
                  :disabled="true"
                />
              </v-col>
              <v-col cols="4">
                <v-text-field
                  v-model="creatforms.port"
                  label="port"
                  type="number"
                  :disabled="true"
                />
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="6">
                <v-text-field
                  v-model="creatforms.microserviceDescription"
                  label="duration"
                  type="number"
                />
              </v-col>
              <v-col cols="6">
                <v-select
                  v-model="creatforms.ptype"
                  :items="ptypes"
                  :rules="[v => !!v || 'Item is required']"
                  label="Type"
                  required
                />
              </v-col>
            </v-row>
            <v-select
              v-model="creatforms.ptype"
              :items="ptypes"
              :rules="[v => !!v || 'Item is required']"
              label="Type"
              required
            />
            <v-text-field
              v-model="creatforms.comment"
              label="Comment"
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
            text="Run"
            variant="plain"
            @click="runtest()"
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
        <v-breadcrumbs :items="['我的项目','服务模块','性能数据']" />
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
              placeholder="Search Type..."
              hide-details
              width="200"
            />
          </td>
          <td>
            <v-text-field
              v-model="comment"
              class="ma-2"
              density="compact"
              placeholder="Search Comment..."
              hide-details
              width="200"
            />
          </td>
          <td>
            <v-btn
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
          icon="mdi-widgets"
          @click="() => router.push({ path: '/user/analysis', query: { profileId: (item as any).item.id } })"
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
    import { GetMicroserviceByPageUser } from '@/api/microservice';
    import { GetProfileByPageUser } from '@/api/profile';
    import router from '@/router';
    import axios from 'axios';
    export default {
      data: () => ({
        itemsPerPage: 10,
        headers: [
          {
            title: 'ID',
            align: 'start',
            sortable: false,
            key: 'id',
          },
          { title: 'Type', key: 'ptype', sortable: false, align: 'start' },
          { title: 'Comment', key: 'comment', sortable: false, align: 'start' },
          { title: 'Create_At', key: 'createdAt', sortable: false, align: 'start' },
          { title: '', key: 'action', sortable: false, align: 'end' },
        ],
        serverItems: [],
        loading: true,
        totalItems: 0,
        ptype: '',
        comment: '',
        id: '',
        name: '',
        search: '',
        router,
        createshow: false,
        creatforms:{
          microserviceName: '',
          ptype: '',
          comment: '',
          microserviceId: 0,
          microserviceToken: '',
          microserviceDescription: '',
          ip: '',
          port: 0,
          duration: 0,
          gc: false,
        },
        ptypes:[
          'cpu',
          'allocs',
          'heap',
          'block',
          'mutex',
          'goroutine',
          'threadcreate',
          'trace',
        ]
      }),
      watch: {
        ptype () {
          this.search = String(Date.now())
        },
        id () {
          this.search = String(Date.now())
        },
        comment () {
          this.search = String(Date.now())
        },
      },
      mounted () {
        this.loadMicroservice()
      },
      methods: {
        async loadItems ({ page, itemsPerPage }: { page: number, itemsPerPage: number}) {
            this.loading = true
            await GetProfileByPageUser({
                profileId: this.id as unknown as number,
                page: page,
                pageSize: itemsPerPage,
                microserviceId: router.currentRoute.value.query.microserviceId as unknown as number,
                ptype: this.ptype,
                comment: this.comment,
            }).then((res) => {
              const data = res.data.data || []
              this.serverItems = data.profiles || []
              this.totalItems = data.total || 0
              this.loading = false
            })
        },
        async loadMicroservice () {
          await GetMicroserviceByPageUser({
            MicroserviceId: router.currentRoute.value.query.microserviceId as unknown as number,
            page: 1,
            pageSize: 1,
            Ip: '',
            MicroserviceDescription: '',
            MicroserviceName: '',
            Port: 0,
            ProjectId: 0
          }).then(res=>{
            const data = res.data.data || []
            this.creatforms.microserviceId = data.microservices[0].id
            this.creatforms.microserviceToken = data.microservices[0].microserviceToken
            this.creatforms.ip = data.microservices[0].ip
            this.creatforms.port = data.microservices[0].port
            this.creatforms.microserviceName = data.microservices[0].microserviceName
            this.creatforms.microserviceDescription = data.microservices[0].microserviceDescription
          })
        }, 
        createprofile () {
          this.createshow = true
        },
        async runtest() {
          axios.create({
            headers: {
              'Content-Type': 'application/json',
            }
          }).post(`http://${this.creatforms.ip}:${this.creatforms.port}/runtest`, {
              serviceToken: this.creatforms.microserviceToken,
              comment: this.creatforms.comment,
              type: this.creatforms.ptype,
            }).then(() => {
            this.loadItems({ page: 1, itemsPerPage: 10});
            this.createshow = false;
          }).catch(error => {
            console.error(error);
            this.createshow = false;
          });
        }
      }
    }
</script>
  
<style scoped>
.v-data-table-container {
    overflow-y: auto;
}
</style>