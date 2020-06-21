<template>
    <div>
        <div class="search-term">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                此处请使用表单生成器生成form填充 表单默认绑定 formData 如手动修改过请自行修改key
                <el-form-item>
                    <el-button @click="openDialog" type="primary">新增</el-button>
                </el-form-item>
            </el-form>
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
            <el-table-column label="日期" width="180">
                <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
            </el-table-column>

            <el-table-column label="姓名" prop="username" width="120"></el-table-column>

            <el-table-column label="性别" prop="gender" width="120"></el-table-column>

            <el-table-column label="出生日期" prop="birthday" width="120"></el-table-column>

            <el-table-column label="手机号" prop="telphone" width="120"></el-table-column>

            <el-table-column label="微信ID" prop="openId" width="120"></el-table-column>

            <el-table-column label="基本信息ID" prop="titBaseinfoId" width="120"></el-table-column>

            <el-table-column label="职业信息ID" prop="titJobInfoId" width="120"></el-table-column>

            <el-table-column label="专业理解ID" prop="titPerfessionalKnowledgeId" width="120"></el-table-column>

            <el-table-column label="行业观点ID" prop="titIndustryPerspectiveId" width="120"></el-table-column>

            <el-table-column label="按钮组">
                <template slot-scope="scope">
                    <el-button @click="updateTitUser(scope.row)" size="small" type="text">变更</el-button>
                    <el-popover placement="top" width="160" v-model="scope.row.visible">
                        <p>确定要删除吗？</p>
                        <div style="text-align: right; margin: 0">
                            <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                            <el-button type="primary" size="mini" @click="deleteTitUser(scope.row)">确定</el-button>
                        </div>
                        <el-button type="text" size="mini" slot="reference">删除</el-button>
                    </el-popover>
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
    </div>
</template>

<script>
    import {
        createTitUser,
        deleteTitUser,
        updateTitUser,
        findTitUser,
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
                visible: false,
                type: "",
                formData: {
                    username: null,
                    gender: null,
                    birthday: null,
                    telphone: null,
                    openId: null,
                    titBaseinfoId: null,
                    titJobInfoId: null,
                    titPerfessionalKnowledgeId: null,
                    titIndustryPerspectiveId: null,
                }
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
            closeDialog() {
                this.dialogFormVisible = false;
                this.formData = {

                    username: null,
                    gender: null,
                    birthday: null,
                    telphone: null,
                    openId: null,
                    titBaseinfoId: null,
                    titJobInfoId: null,
                    titPerfessionalKnowledgeId: null,
                    titIndustryPerspectiveId: null,
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
</style>
