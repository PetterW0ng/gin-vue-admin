<template>
    <div>
        <el-table
                :data="tableData"
                border
                ref="multipleTable"
                stripe
                style="width: 100%"
                tooltip-effect="dark">
            <el-table-column type="selection" width="55"></el-table-column>
            <el-table-column label="姓名" prop="userName" width="80"></el-table-column>
            <el-table-column label="证书编号" prop="serialNum" width="200"></el-table-column>
            <el-table-column label="手机" prop="phone" width="120"></el-table-column>
            <el-table-column label="发证时间" prop="issueTime" width="100"></el-table-column>
            <el-table-column label="身份证号" prop="idCard" width="180"></el-table-column>
            <el-table-column label="创建日期" width="160">
                <template slot-scope="scope">{{ scope.row.createTime|formatDate }}</template>
            </el-table-column>
            <el-table-column label="证书名称" prop="certificateName" ></el-table-column>
            <el-table-column label="发证单位" prop="issuingUnit"></el-table-column>
            <el-table-column label="按钮组" width="150">
                <template slot-scope="scope">
                    <el-button @click="updateHolder(scope.row)" size="small" type="text">查看证书</el-button>
                    <el-button @click="updateHolder(scope.row)" size="small" type="text">修改</el-button>
                    <el-popover placement="top" width="160" v-model="scope.row.visible">
                        <p>确定要删除吗？</p>
                        <div style="text-align: right; margin: 0">
                            <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                            <el-button type="primary" size="mini" @click="deleteHolder(scope.row)">确定</el-button>
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

        <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新增Api">
            <el-form :inline="true" :model="form" label-width="80px">
                <el-form-item label="客户名">
                    <el-input autocomplete="off" v-model="form.customerName"></el-input>
                </el-form-item>
                <el-form-item label="客户电话">
                    <el-input autocomplete="off" v-model="form.customerPhoneData"></el-input>
                </el-form-item>
            </el-form>
            <div class="dialog-footer" slot="footer">
                <el-button @click="closeDialog">取 消</el-button>
                <el-button @click="enterDialog" type="primary">确 定</el-button>
            </div>
        </el-dialog>
        <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="查看证书图片">
            <el-form :inline="true" :model="form" label-width="500px">
                <img :src="ss"/>
            </el-form>
            <div class="dialog-footer" slot="footer">
                <el-button @click="closeImgViewDialog">关 闭</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
    import {
        getCertHolderList,deleteCertHolder,getCertHolder,updateCertHolder
    } from "@/api/certHolder";
    import { formatTimeToStr } from "@/utils/data";
    import infoList from "@/components/mixins/infoList";
    export default {
        name:"holder",
        mixins: [infoList],
        data() {
            return {
                listApi: getCertHolderList,
                dialogFormVisible: false,
                imgViewDialog: false,
                visible: false,
                type: "",
                form: {
                    customerName: "",
                    customerPhoneData: ""
                }
            }
        },
        filters: {
            formatDate: function(time) {
                if (time != null && time != "") {
                    let date = new Date(time*1000);
                    return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
                } else {
                    return "";
                }
            }
        },
        methods:{
            async updateHolder(row) {
                const res = await getCertHolder({ ID: row.ID });
                this.type = "update";
                if (res.code == 0) {
                    this.form = res.data.customer;
                    this.dialogFormVisible = true;
                }
            },
            closeDialog() {
                this.dialogFormVisible = false;
                this.form = {
                    customerName: "",
                    customerPhoneData: ""
                };
            },
            closeImgViewDialog() {
                this.imgViewDialog = false;
            },
            async deleteHolder(row) {
                this.visible = false;
                const res = await deleteCertHolder({ID: row.ID});
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
                        break;
                    case "update":
                        res = await updateCertHolder(this.form);
                        break;
                    default:
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
            this.searchInfo.userName="";
            this.getTableData();
        }
    }


</script>
<style lang="scss">

</style>
