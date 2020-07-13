<template>
 <div class="big">
     <div class="mid">
         <!--<el-row :gutter="32">
             <el-col :xs="24" :sm="24" :lg="8">
                 <div class="chart-wrapper">
                     <raddar-chart />
                 </div>
             </el-col>
             <el-col :xs="24" :sm="24" :lg="8">
                 <div class="chart-wrapper">
                     <stackMap />
                 </div>
             </el-col>
             <el-col :xs="24" :sm="24" :lg="8">
                 <div class="chart-wrapper">
                     <Sunburst/>
                 </div>
             </el-col>
         </el-row>-->
         正在开发中 ...
     </div>
     <div class="top">
         <!--         <div id="main" class="chart-container"></div>-->
     </div>
      <div class="bottom">
          <!--<el-row :gutter="32">
              <el-col :xs="24" :sm="24" :lg="12">
                  <div class="chart-player">
                      <musicPlayer />
                  </div>
              </el-col>
              <el-col :xs="24" :sm="24" :lg="12">
                  <div  class="chart-player">
                      <todo-list />
                  </div>
              </el-col>
          </el-row>-->
     </div>
 </div>

</template>

<script>
    // import echarts from 'echarts'
    // require('echarts/theme/macarons') // echarts theme
    // import RaddarChart from "./component/RaddarChart"
    // import stackMap from "./component/stackMap"
    // import Sunburst from "./component/Sunburst"
    // import musicPlayer from "./component/musicPlayer"
    // import TodoList from "./component/todoList"

    export default {
        name: 'Dashboard',
        data() {
            return {}
        },
        // components:{
        //     RaddarChart, //雷达图
        //     stackMap, //堆叠图
        //     Sunburst, //旭日图
        //     musicPlayer,  //音乐播放器
        //     TodoList //TodoList
        // },
        mounted() {

        }
    }
</script>

<style lang="scss" scoped>
    .big{
        margin:100px 0 0 0;
        padding-top: 0;
        background-color: rgb(243,243,243);;
        .top{
            width: 100%;
            height: 360px;
            margin-top: 20px;
            overflow: hidden;
            .chart-container{
                position: relative;
                width: 100%;
                height: 100%;
                padding: 20px;
                background-color: #fff;
            }
        }
        .mid{
            width: 100%;
            height: 380px;
            .chart-wrapper {
                height: 340px;
                background: #fff;
                padding: 16px 16px 0;
                margin-bottom: 32px;
            }
        }
        .bottom{
            width: 100%;
            height: 300px;
            margin: 20px 0;
            .el-row{
                margin-right: 4px !important;
            }
            .chart-player{
                width: 100%;
                height: 270px;
                padding: 10px;
                background-color: #fff;
            }
        }
    }

</style>
