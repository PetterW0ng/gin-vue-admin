<template>
    <div>
        <div class="search-term">
            <!--<el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                此处请使用表单生成器生成form填充 表单默认绑定 formData 如手动修改过请自行修改key
                <el-form-item>
                    <el-button @click="openDialog" type="primary">新增</el-button>
                </el-form-item>
            </el-form>-->
        </div>
        <el-table
                :data="tableData"
                border
                ref="multipleTable"
                stripe
                style="width: 100%"
                tooltip-effect="dark"
        >
            <el-table-column type="selection" width="55"></el-table-column>
            <el-table-column label="注册日期" width="180">
                <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
            </el-table-column>

            <el-table-column label="姓名" prop="username" width="120"></el-table-column>

            <el-table-column label="性别" prop="gender" width="120" :formatter="genderFormat"></el-table-column>

            <el-table-column label="出生日期" prop="birthday" width="120" :formatter="birthdayFormat"></el-table-column>

            <el-table-column label="手机号" prop="telphone" width="120"></el-table-column>

            <el-table-column label="换工作意向" prop="changeJobOption" width="120"
                             :formatter="jobInfoFormat"></el-table-column>

            <el-table-column label="按钮组">
                <template slot-scope="scope">
                    <el-popover placement="top" width="160" v-model="scope.row.visible">
                        <p>确定要删除吗？</p>
                        <div style="text-align: right; margin: 0">
                            <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                            <el-button type="primary" size="mini" @click="deleteTitUser(scope.row)">确定</el-button>
                        </div>
                        <el-button type="text" size="mini" slot="reference">删除</el-button>
                    </el-popover>
                    <el-button @click="viewBaseinfo(scope.row)" size="small" type="text">基本信息</el-button>
                </template>
            </el-table-column>
        </el-table>

        <el-pagination
                :current-page="page"
                :page-size="pageSize"
                :page-sizes="[10, 30, 50, 100]"
                :style="{float:'right',padding:'20px'}"
                :total="total"
                @current-change="handleCurrentChange"
                @size-change="handleSizeChange"
                layout="total, sizes, prev, pager, next, jumper"
        ></el-pagination>

        <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
            此处请使用表单生成器生成form填充 表单默认绑定 formData 如手动修改过请自行修改key
            <div class="dialog-footer" slot="footer">
                <el-button @click="closeDialog">取 消</el-button>
                <el-button @click="enterDialog" type="primary">确 定</el-button>
            </div>
        </el-dialog>
        <!-- 基本信息 -->
        <el-dialog :visible.sync="dialogBaseinfoVisible" title="基本信息">
            <el-form label-position="right" label-width="400px" :model="userBaseinfo">
                <!-- 教育背景 -->
                <el-form-item label="教育背景" class="baseInfoItem"/>
                <el-form-item label="毕业学校" class="baseInfoItem">
                    <el-input v-model="userBaseinfo.school" readonly></el-input>
                </el-form-item>
                <el-form-item label="所修专业" class="baseInfoItem">
                    <el-input v-model="userBaseinfo.majorsStudied" readonly></el-input>
                </el-form-item>
                <el-form-item label="最高学历" class="baseInfoItem">
                    <el-input v-model="userBaseinfo.highestEducation" readonly></el-input>
                </el-form-item>
                <el-form-item label="所上学制" class="baseInfoItem">
                    <el-input v-model="userBaseinfo.schoolSystem" readonly></el-input>
                </el-form-item>
                <el-form-item label="工作经历" class="baseInfoItem"/>
                <el-form-item label="工作状态" class="baseInfoItem">
                    <el-input v-model="userBaseinfo.workingState" readonly></el-input>
                </el-form-item>
                <el-form-item label="工作单位" class="baseInfoItem">
                    <el-input v-model="userBaseinfo.company" readonly></el-input>
                </el-form-item>
                <el-form-item label="单位地域" class="baseInfoItem">
                    <el-input v-model="userBaseinfo.area" readonly></el-input>
                </el-form-item>
                <el-form-item label="职务职位" class="baseInfoItem">
                    <el-input v-model="userBaseinfo.jobTitle" readonly></el-input>
                </el-form-item>
                <el-form-item label="所在单位服务模式" class="baseInfoItem">
                    <el-input v-model="userBaseinfo.serviceType" readonly></el-input>
                </el-form-item>
                <el-form-item label="我的月度收入水平" class="baseInfoItem">
                    <el-input v-model="userBaseinfo.income" readonly></el-input>
                </el-form-item>
                <el-form-item label="我的福利待遇情况" class="baseInfoItem">
                    <el-input v-model="userBaseinfo.benefits" readonly></el-input>
                </el-form-item>
                <el-form-item label="我目前服务的儿童类型" class="baseInfoItem">
                    <el-input v-model="userBaseinfo.childType" readonly></el-input>
                </el-form-item>
                <el-form-item label="我目前主要服务儿童年龄段" class="baseInfoItem">
                    <el-input v-model="userBaseinfo.childAge" readonly></el-input>
                </el-form-item>
                <!-- 培训经历 -->
                <el-form-item label="培训经历" class="baseInfoItem"/>
                <el-form-item label="我到目前为止参加培训数量（一周内的算作短期培训）" class="baseInfoItem">
                    <el-input v-model="userBaseinfo.trainingNumber" readonly></el-input>
                </el-form-item>
                <el-form-item label="我去年在培训上的花费" class="baseInfoItem">
                    <el-input v-model="userBaseinfo.trainingFee" readonly></el-input>
                </el-form-item>
                <el-form-item label="我以往培训详情" class="baseInfoItem"/>
                <div v-for="(item, index) in userBaseinfo.trainingInfos" :key="index"
                     :value="'培训'+(index + 1)">
                    <el-form-item label="参训课程" class="baseInfoItem">
                        <el-input v-model="item.trainingCourse" readonly></el-input>
                    </el-form-item>
                    <el-form-item label="开始时间" class="baseInfoItem">
                        <el-input :value="item.beginTime | formatDate" readonly></el-input>
                    </el-form-item>
                    <el-form-item label="结束时间" class="baseInfoItem">
                        <el-input :value="item.endTime | formatDate" readonly></el-input>
                    </el-form-item>
                    <el-form-item label="付费方式" class="baseInfoItem">
                        <el-input v-model="item.paymentWay" readonly></el-input>
                    </el-form-item>
                </div>
            </el-form>
        </el-dialog>
    </div>
</template>

<script>
    import {
        createTitUser,
        deleteTitUser,
        updateTitUser,
        findTitUser,
        findBaseInfo,
        getTitUserList
    } from "@/api/tit_user";  //  此处请自行替换地址
    import {formatTimeToStr} from "@/utils/data";
    import infoList from "@/components/mixins/infoList";

    export default {
        name: "TitUser",
        mixins: [infoList],
        data() {
            return {
                listApi: getTitUserList,
                dialogFormVisible: false,
                dialogBaseinfoVisible: false,
                visible: false,
                type: "",
                formData: {
                    username: null,
                    gender: null,
                    birthday: null,
                    telphone: null,
                    changeJobOption: null,
                },
                userBaseinfo: {}
            };
        },
        filters: {
            formatDate: function (time) {
                if (time != null && time != "") {
                    var date = new Date(time);
                    return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
                } else {
                    return "";
                }
            }
        },
        methods: {

            async updateTitUser(row) {
                const res = await findTitUser({ID: row.ID});
                this.type = "update";
                if (res.code == 0) {
                    this.formData = res.data.reuser;
                    this.dialogFormVisible = true;
                }
            },
            async viewBaseinfo(row) {
                const res = await findBaseInfo({ID: row.ID});
                if (res.code == 0) {
                    this.userBaseinfo = res.data.userBaseinfo;
                    this.dialogBaseinfoVisible = true;
                }
            },
            genderFormat(row) {
                return row.gender == 1 ? "男" : "女"
            },
            jobInfoFormat(row) {
                return row.changeJobOption == 0 ? "未知" : (row.changeJobOption == 61 ? "不考虑" : (row.changeJobOption == 62 ? "考虑换其他单位" : "考虑换其他行业"))
            },
            birthdayFormat(row) {
                return formatTimeToStr(row.birthday, "yyyy-MM-dd");
            },
            closeDialog() {
                this.dialogFormVisible = false;
                this.formData = {
                    username: null,
                    gender: null,
                    birthday: null,
                    telphone: null,
                };
            },
            async deleteTitUser(row) {
                this.visible = false;
                const res = await deleteTitUser({ID: row.ID});
                if (res.code == 0) {
                    this.$message({
                        type: "success",
                        message: "删除成功"
                    });
                    this.getTableData();
                }
            },
            async enterDialog() {
                let res;
                switch (this.type) {
                    case "create":
                        res = await createTitUser(this.formData);
                        break;
                    case "update":
                        res = await updateTitUser(this.formData);
                        break;
                    default:
                        res = await createTitUser(this.formData);
                        break;
                }

                if (res.code == 0) {
                    this.closeDialog();
                    this.getTableData();
                }
            },
            openDialog() {
                this.type = "create";
                this.dialogFormVisible = true;
            }
        },
        created() {
            this.getTableData();
        }
    };
</script>

<style>
    .el-table .warning-row {
        background: orange;
    }

    .el-table .success-row {
        background: green;
    }

    .el-table .error-row {
        background: red;
    }

    .baseInfoItem {
        padding-bottom: 0px;
        padding-top: 0px;
    }
</style>
