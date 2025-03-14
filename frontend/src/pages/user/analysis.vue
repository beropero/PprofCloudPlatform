<template>
  <v-app
    class="v-data-table-container"
    style="background-color: white;color: black;overflow: auto;"
  >
    <v-breadcrumbs
      :items="['我的项目','服务模块','性能数据','数据分析']"
      style="font-size: small;"
    />
    <v-sheet
      style="background-color: white;color: black;"
    >
      <v-textarea
        v-model="pproftitle"
        label="pprof"
        hide-details
        density="compact"
        rows="4"
        variant="outlined"
        style="background-color: white;color: black;margin: 8px;"
        :readonly="true"
      />
    </v-sheet>
    <FlameChart :profile="pprof" />
    <v-data-table-server
      v-model:items-per-page="itemsPerPage"  
      :headers="headers as any" 
      :items="serverItems"
      :items-length="tabledata.length"
      style="background-color: white;color: black;"
      :loading="loading"
      @update:options="loadItems"
    >
      <template #top>
        <tr>
          <td>
            <v-select
              v-model="selectedItem"
              label="Select"
              :items="['Samples', 'Locations', 'Mappings']"
              class="ma-2"
              density="compact"
              width="200px"
              hide-details
              variant="outlined"
              @update:model-value="selectedItemChanged"
            />
          </td>
          
          <v-divider />
        </tr>
      </template>
    </v-data-table-server>
  </v-app>
</template>

<script setup lang="ts">
import { GetProfileContent } from '@/api/profile';
import router from '@/router';
import { Profile } from '@/utils/pprof';
import FlameChart from '@/components/FlameChart.vue';
const pprof = ref();
const pprofstring = ref('');
const pproftitle = ref('')
const tabledata = ref([]);
const itemsPerPage = ref(8);
const loading = ref(false);
const selectedItem = ref('Samples');
const serverItems = ref([]);

const loadItems = ({ page, itemsPerPage }: { page: number, itemsPerPage: number }) => {
  loading.value = true;
  const start = (page - 1) * itemsPerPage;
  const end = start + itemsPerPage;
  serverItems.value = tabledata.value.slice(start, end);
  loading.value = false;
  console.log(tabledata.value);
};

const selectedItemChanged = (value: string) => {
  if (value === 'Locations') {
    headers.value = [
          {
            title: 'ID',
            align: 'start',
            sortable: false,
            key: '1',
          },
          { title: 'Address', key: '2', sortable: false, align: 'start' },
          { title: 'M', key: '3', sortable: false, align: 'start' },
          { title: 'Function', key: '4', sortable: false, align: 'start' },
          { title: 'Location', key: '5', sortable: false, align: 'start' },
          { title: 'SystemName', key: '6', sortable: false, align: 'start' },
        ]
        tabledata.value = pprof.value.getLocationItems()
  }
  if (value === 'Samples') {
    headers.value = [];
    pprof.value.getSampleHeader().forEach((item:string, index:number) => {
     headers.value.push({
       title: item,
       key: index.toString(),
       sortable: false,
       align: 'start',
     })
   });
   headers.value.push({
       title: '',
       key: headers.value.length.toString(),
       sortable: false,
       align: 'start',
   },
   {
       title: '',
       key: (headers.value.length+1).toString(),
       sortable: false,
       align: 'start',
   })
   tabledata.value = pprof.value.getSampleItems()
  }
  if (value === 'Mappings') {
    headers.value = [
          {
            title: 'ID',
            align: 'start',
            sortable: false,
            key: '1',
          },
          { title: 'Start', key: '2', sortable: false, align: 'start' },
          { title: 'Offset', key: '3', sortable: false, align: 'start' },
          { title: 'Limit',key: '4', sortable: false, align: 'start' },
          { title: 'File', key: '5', sortable: false, align: 'start' },
          { title: 'BuildID', key: '6', sortable: false, align: 'start' },
          { title: 'bits', key: '7', sortable: false, align: 'start' },
    ]
    tabledata.value = pprof.value.getMappingItems()
  }
  loadItems({ page: 1, itemsPerPage: 8 });
};
interface Header {
  title: string;
  key: string;
  sortable: boolean;
  align: string;
}
const headers = ref<Header[]>([]);
onMounted(() => {
 GetProfileContent({
    id : router.currentRoute.value.query.profileId as unknown as string
 }).then(res => {
   const pprofdata = res.data.data.fileContent
   pprof.value = new Profile(pprofdata)
   console.log(pprofdata)
   pprofstring.value = pprof.value.toString()
   pprof.value.getSampleHeader().forEach((item:string, index:number) => {
     headers.value.push({
       title: item,
       key: index.toString(),
       sortable: false,
       align: 'start',
     })
   });
   headers.value.push({
       title: '',
       key: headers.value.length.toString(),
       sortable: false,
       align: 'start',
   },
   {
       title: '',
       key: (headers.value.length+1).toString(),
       sortable: false,
       align: 'start',
   })
   pproftitle.value = pprof.value.getTitile()
   tabledata.value = pprof.value.getSampleItems()
   loadItems({ page: 1, itemsPerPage: itemsPerPage.value });
 })
 loading.value = false
})
</script>

<style scoped>
.text-overflow {
  white-space: pre-wrap;
  text-overflow: ellipsis;
  max-width: 200px; 
}
</style>