<template>
    <div>
        <van-nav-bar title="完善基本信息"
                     :fixed=true
                     :border=false
                     style="height: 2.5rem;padding-top:0.5rem"/>
        <div class="mainDiv">
            <!-- 教育背景 -->
            <van-cell-group v-show="currentStep == 1">
                <van-cell title="" value="" class="title" label="1. 教育背景"/>
                <van-cell title="请填写最高学历信息" class="tips"/>
                <van-field v-model="baseinfo.school"
                           label="毕业学校"
                           required/>
                <van-field :value="majorsStudied"
                           label="所修专业" is-link readonly
                           @click="showMajorsStudiedView = true"
                           required/>
                <!--<el-form>
                    <el-form-item label="所修专业" required style="margin: 0px 10px 0px 7px;padding-right: 10px">
                        <el-select v-model="baseinfo.majorsStudied" filterable clearable placeholder="所修专业"
                                   style="margin-left:20px;width: 230px;">
                            <el-option
                                    v-for="item in majorsStudiedColumns"
                                    :key="item.propertyValue"
                                    :label="item.text"
                                    :value="item.propertyValue">
                            </el-option>
                        </el-select>
                    </el-form-item>
                </el-form>-->
                <van-field :value="highestEducation"
                           label="最高学历"
                           required is-link
                           @click="showEducationPopView = true"
                           readonly/>
                <van-field :value="schoolSystem"
                           label="所上学制"
                           required is-link
                           @click="showSchoolSystemPopView = true"
                           readonly/>
            </van-cell-group>
            <!-- 工作单位 -->
            <van-cell-group v-show="currentStep == 2">
                <van-cell title="" value="" class="title" label="2. 工作单位"/>
                <van-cell title="请填写目前工作情况" class="tips"/>
                <van-field name="radio" required label="工作状态">
                    <template #input>
                        <van-radio-group v-model="baseinfo.workingState" @change="workingStateChangeHandle"
                                         direction="horizontal">
                            <van-radio v-for="(item, index) in workingStateOption" style="width: 100px"
                                       :name="item.propertyValue"
                                       :value="item.propertyValue" :key="item.propertyValue">{{item.text}}
                            </van-radio>
                        </van-radio-group>
                    </template>
                </van-field>
                <van-field v-model="baseinfo.company" :disabled="disabled" v-if="!disabled"
                           label="工作单位"
                           required/>
                <van-field v-model="areaSelected" :disabled="disabled" v-if="!disabled"
                           label="单位地域"
                           @click="showAreaPopViewH"
                           required readonly is-link/>
                <van-field v-model="baseinfo.jobTitle" :disabled="disabled" v-if="!disabled"
                           label="职务职位" required/>
                <van-field v-model="serviceType" :disabled="disabled" v-if="!disabled"
                           label="所在单位服务模式"
                           @click="showServiceTypePopViewH"
                           required readonly is-link/>
                <van-field v-model="income" :disabled="disabled" v-if="!disabled"
                           label="我的月度收入水平"
                           @click="showIncomePopViewH"
                           required readonly is-link/>
                <van-field name="benefits" label="我的福利待遇情况" required v-if="!disabled">
                    <template #input>
                        <van-checkbox-group v-model="baseinfo.benefits" :disabled="disabled" direction="horizontal">
                            <van-checkbox v-for="(item, index) in benefitsOption" :name="item.propertyValue"
                                          :value="item.propertyValue" :key="item.propertyValue" shape="square">
                                {{item.text}}
                            </van-checkbox>
                        </van-checkbox-group>
                    </template>
                </van-field>
                <van-field name="childType" label="我目前服务的儿童类型" required>
                    <template #input>
                        <van-checkbox-group v-model="baseinfo.childType" direction="horizontal">
                            <van-checkbox v-for="(item, index) in childTypeOption" :name="item.propertyValue"
                                          :value="item.propertyValue" :key="item.propertyValue" shape="square">
                                {{item.text}}
                            </van-checkbox>
                        </van-checkbox-group>
                    </template>
                </van-field>
                <van-field v-model="childAge"
                           label="我目前主要服务儿童年龄段"
                           @click="showChildAgePopView = true"
                           required readonly is-link/>
            </van-cell-group>
            <!-- 培训经历 -->
            <van-cell-group v-show="currentStep == 3">
                <van-cell title="" value="" class="title" label="3.培训经历"/>
                <van-field required label="我到目前为止参加培训数量（一周内的算作短期培训）">
                    <template #input>
                        <van-checkbox-group v-model="baseinfo.trainingNumber" direction="horizontal">
                            <van-checkbox v-for="(item, index) in trainingNumberOption" :name="item.propertyValue"
                                          :value="item.propertyValue" :key="item.propertyValue" shape="square">
                                {{item.text}}
                            </van-checkbox>
                        </van-checkbox-group>
                    </template>
                </van-field>
                <van-field v-model="trainingFee"
                           label="我去年在培训上的花费"
                           @click="showTrainingFeePopView = true"
                           required readonly is-link/>
                <van-cell title="我以往培训详情，请描述"/>
                <van-cell-group v-for="(item, index) in trainingInfos" :key="index">
                    <van-cell :title="'培训'+(index + 1)" class="px-title" icon="shop-o">
                        <template #right-icon>
                            <van-icon name="clear" style="line-height: inherit;" @click="removeTrainingInfo(index)"/>
                        </template>
                    </van-cell>
                    <van-field v-model="item.trainingCourse"
                               label="培训课程" required/>
                    <van-field v-model="item.beginTimeStr" is-link readonly @click="handleSelectTime(1, index)"
                               label="开始时间" required/>
                    <van-field v-model="item.endTimeStr" is-link readonly @click="handleSelectTime(2, index)"
                               label="结束时间" required/>
                    <van-field v-model="item.paymentWayStr"
                               label="付费方式"
                               @click="handlePaymentWay(index)"
                               readonly is-link required/>
                </van-cell-group>
                <van-button plain hairline type="info" size="large" @click="addTrainingInfo">添加培训经历</van-button>
            </van-cell-group>
        </div>
        <div class="buttonDiv">
            <van-button class="button" type="info" size="large" @click="saveBaseinfo" v-if="currentStep == 3">提 交
            </van-button>
            <van-button class="button" type="info" size="large" @click="next" v-if="currentStep < 3">下一步</van-button>
            <van-button class="button" type="info" size="large" @click="previous" v-if="currentStep > 1">上一步
            </van-button>
        </div>
        <!-- 弹框选项开始 -->
        <van-popup v-model="showSchoolPopView" round position="bottom">
            <van-picker show-toolbar title=""
                        :columns="schoolColumns"
                        @cancel="showSchoolPopView = false"
                        @confirm="onConfirmSchool"/>
        </van-popup>

        <!-- 时间选择器 -->
        <van-popup v-model="showDateTimePopView" round position="bottom">
            <van-datetime-picker v-model="currentDate" type="date" @confirm="onConfirmDateTime"
                                 @cancel="showDateTimePopView = false"
                                 :formatter="formatter"
                                 :max-date="maxDate"
                                 :min-date="minDate"/>
        </van-popup>

        <van-popup v-model="showMajorsStudiedView" round position="bottom">
            <van-picker show-toolbar title=""
                        :columns="majorsStudiedColumns"
                        @cancel="showMajorsStudiedView = false"
                        @confirm="onConfirmMajorsStudied"/>
        </van-popup>

        <van-popup v-model="showTrainingFeePopView" round position="bottom">
            <van-picker show-toolbar title=""
                        :columns="trainingFeeColumns"
                        @cancel="showTrainingFeePopView = false"
                        @confirm="onConfirmTrainingFee"/>
        </van-popup>

        <van-popup v-model="showChildAgePopView" round position="bottom">
            <van-picker show-toolbar title=""
                        :columns="childAgeColumns"
                        @cancel="showChildAgePopView = false"
                        @confirm="onConfirmChildAge"/>
        </van-popup>

        <van-popup v-model="showIncomePopView" round position="bottom">
            <van-picker show-toolbar title=""
                        :columns="incomeColumns"
                        @cancel="showIncomePopView = false"
                        @confirm="onConfirmIncome"/>
        </van-popup>

        <van-popup v-model="showServiceTypePopView" round position="bottom">
            <van-picker show-toolbar title=""
                        :columns="serviceTypeColumns"
                        @cancel="showServiceTypePopView = false"
                        @confirm="onConfirmServiceType"/>
        </van-popup>

        <van-popup v-model="showAreaPopView" round position="bottom">
            <!--<van-picker show-toolbar  title=""
                    :columns="areaColumns"
                    @cancel="showAreaPopView = false"
                    @confirm="onConfirmArea" />-->
            <van-area :area-list="areas" ref="areaSelectPicker" id="areaSelectPicker" title="" :value="baseinfo.area"
                      @cancel="showAreaPopView = false"
                      @confirm="onConfirmArea"/>
        </van-popup>

        <van-popup v-model="showSchoolSystemPopView" round position="bottom">
            <van-picker show-toolbar title=""
                        :columns="schoolSystemColumns"
                        @cancel="showSchoolSystemPopView = false"
                        @confirm="onConfirmSchoolSystem"/>
        </van-popup>

        <van-popup v-model="showEducationPopView" round position="bottom">
            <van-picker show-toolbar title=""
                        :columns="educationColumns"
                        @cancel="showEducationPopView = false"
                        @confirm="onConfirmEducation"/>
        </van-popup>

        <van-popup v-model="showPaymentWayPopView" round position="bottom">
            <van-picker show-toolbar title=""
                        :columns="paymentWayColumns"
                        @cancel="showPaymentWayPopView = false"
                        @confirm="onConfirmPaymentWay"/>
        </van-popup>
        <!-- 弹框选项结束 -->

    </div>
</template>

<script>
    import {mapState} from "vuex";
    import {selectOptions, getArea, saveBaseinfo, initUserBaseinfo} from "../../../serve/api";
    import areas from '../../../config/area.js'
    import {Toast} from "vant";
    import Moment from "moment";

    export default {
        data() {
            return {
                disabled: true,
                showSchoolPopView: false,
                schoolColumns: [],

                showMajorsStudiedView: false,
                majorsStudiedColumns: [],

                showTrainingFeePopView: false,
                trainingFeeColumns: [],

                showChildAgePopView: false,
                childAgeColumns: [],

                showIncomePopView: false,
                incomeColumns: [],

                showServiceTypePopView: false,
                serviceTypeColumns: [],

                showAreaPopView: false,
                areaColumns: [],

                showSchoolSystemPopView: false,
                schoolSystemColumns: [],

                showEducationPopView: false,
                educationColumns: [],

                showPaymentWayPopView: false,
                paymentWayNum: 0,
                paymentWayColumns: [],

                showDateTimePopView: false,
                currentDate: new Date('2000/01/01'),
                minDate: new Date('1949/01/01'),
                maxDate: new Date(),
                //1:begin  2:end
                timePickerType: 0,
                timePickerNum: 0,

                trainingNumberOption: [],
                benefitsOption: [],
                childTypeOption: [],
                workingStateOption: [],
                currentStep: 1,
                areas: areas,
                areaSelected: '',
                modify: -1,
                // 基本信息
                baseinfo: {
                    // 教育背景
                    school: '',
                    majorsStudied: 0,
                    highestEducation: 0,
                    schoolSystem: 0,
                    // 工作单位
                    workingState: 0,
                    company: '',
                    area: '',
                    jobTitle: '',
                    serviceType: 0,
                    income: 0,
                    benefits: [],
                    childType: [],
                    childAge: 0,
                    // 培训经历
                    trainingNumber: [],
                    trainingFee: 0,
                    trainingInfos: []
                },
            }
        },
        computed: {
            ...mapState(['userInfo']),
            school() {
                return this.getPropertyName(this.schoolColumns, this.baseinfo.school)
            },
            majorsStudied() {
                return this.getPropertyName(this.majorsStudiedColumns, this.baseinfo.majorsStudied)
            },
            trainingFee() {
                return this.getPropertyName(this.trainingFeeColumns, this.baseinfo.trainingFee)
            },
            childAge() {
                return this.getPropertyName(this.childAgeColumns, this.baseinfo.childAge)
            },
            income() {
                return this.getPropertyName(this.incomeColumns, this.baseinfo.income)
            },
            serviceType() {
                return this.getPropertyName(this.serviceTypeColumns, this.baseinfo.serviceType)
            },
            schoolSystem() {
                return this.getPropertyName(this.schoolSystemColumns, this.baseinfo.schoolSystem)
            },
            highestEducation() {
                return this.getPropertyName(this.educationColumns, this.baseinfo.highestEducation)
            },
            trainingInfos() {
                let newArr = [...this.baseinfo.trainingInfos]
                newArr.map(el => {
                    return el.beginTimeStr = el.beginTime ? Moment(el.beginTime).format('YYYY-MM-DD') : "", el.endTimeStr = el.endTime ? Moment(el.endTime).format('YYYY-MM-DD') : "", el.paymentWayStr = this.getPropertyName(this.paymentWayColumns, el.paymentWay)
                })
                return newArr
            }
        },
        created() {
            this.modify = this.$route.query.modify
            this.initSelectOptions()
        },
        mounted() {
        },
        methods: {
            workingStateChangeHandle(value) {
                if (value == 2) {
                    this.disabled = true
                } else {
                    this.disabled = false
                }
            },
            showAreaPopViewH() {
                if (!this.disabled) {
                    this.showAreaPopView = true
                }
            },
            showIncomePopViewH() {
                if (!this.disabled) {
                    this.showIncomePopView = true
                }
            },
            showServiceTypePopViewH() {
                if (!this.disabled) {
                    this.showServiceTypePopView = true
                }
            },
            previous() {
                this.currentStep -= 1
            },
            next() {
                let validated = true
                if (this.currentStep == 1) {
                    if (this.baseinfo.school.length == 0) {
                        validated = false
                        Toast.fail("毕业学校为必填项!")
                    } else if (this.baseinfo.majorsStudied == 0) {
                        validated = false
                        Toast.fail("所修专业为必填项!")
                    } else if (this.baseinfo.highestEducation == 0) {
                        validated = false
                        Toast.fail("最高学历为必填项!")
                    } else if (this.baseinfo.schoolSystem == 0) {
                        validated = false
                        Toast.fail("所上学制为必填项!")
                    }
                } else if (this.currentStep == 2) {
                    if (this.baseinfo.workingState == 0) {
                        validated = false
                        Toast.fail("工作状态为必填项!")
                    } else if (this.baseinfo.workingState == 1) {
                        if (this.baseinfo.company.length == 0) {
                            validated = false
                            Toast.fail("工作单位为必填项!")
                        } else if (this.baseinfo.area.length == 0) {
                            validated = false
                            Toast.fail("单位地域为必填项!")
                        } else if (this.baseinfo.jobTitle.length == 0) {
                            validated = false
                            Toast.fail("职务职位为必填项!")
                        } else if (this.baseinfo.serviceType == 0) {
                            validated = false
                            Toast.fail("服务模式为必填项!")
                        } else if (this.baseinfo.income == 0) {
                            validated = false
                            Toast.fail("月度收入水平为必填项!")
                        } else if (this.baseinfo.benefits.length == 0) {
                            validated = false
                            Toast.fail("福利待遇情况为必填项!")
                        }
                    }
                    if (this.baseinfo.childType.length == 0) {
                        validated = false
                        Toast.fail("儿童类型为必填项!")
                    } else if (this.baseinfo.childAge == 0) {
                        validated = false
                        Toast.fail("儿童年龄段为必填项!")
                    }
                } else {
                }
                if (validated) {
                    this.currentStep += 1
                }
            },
            removeTrainingInfo(index) {
                this.baseinfo.trainingInfos.splice(index, 1)
            },
            addTrainingInfo() {
                this.baseinfo.trainingInfos.push({})
            },
            handlePaymentWay(index) {
                this.showPaymentWayPopView = true
                this.paymentWayNum = index
            },
            onConfirmPaymentWay(value) {
                this.showPaymentWayPopView = false
                let trainingInfo = this.baseinfo.trainingInfos[this.paymentWayNum]
                trainingInfo.paymentWay = value.propertyValue;
                this.$set(this.baseinfo.trainingInfos, this.paymentWayNum, trainingInfo);
            },
            // 格式化DateTime pickView
            formatter(type, value) {
                if (type === 'year') {
                    return `${value}`;
                } else if (type === 'month') {
                    return `${value}`
                } else if (type === 'day') {
                    return `${value}`
                }
                return value;
            },
            handleSelectTime(bOe, index) {
                this.showDateTimePopView = true
                this.timePickerNum = index
                this.timePickerType = bOe
            },
            onConfirmDateTime(value) {
                let trainingInfo = this.baseinfo.trainingInfos[this.timePickerNum]
                if (this.timePickerType == 1) {
                    trainingInfo.beginTime = Moment(value).format();
                } else {
                    trainingInfo.endTime = Moment(value).format();
                }
                this.$set(this.baseinfo.trainingInfos, this.timePickerNum, trainingInfo);
                this.showDateTimePopView = false;
            },
            onConfirmSchool(value, index) {
                this.baseinfo.school = value.propertyValue
                this.showSchoolPopView = false;
            },
            onConfirmMajorsStudied(value, index) {
                this.baseinfo.majorsStudied = value.propertyValue
                this.showMajorsStudiedView = false;
            },
            onConfirmTrainingFee(value, index) {
                this.baseinfo.trainingFee = value.propertyValue
                this.showTrainingFeePopView = false;
            },
            onConfirmChildAge(value, index) {
                this.baseinfo.childAge = value.propertyValue
                this.showChildAgePopView = false;
            },
            onConfirmIncome(value, index) {
                this.baseinfo.income = value.propertyValue
                this.showIncomePopView = false;
            },
            onConfirmServiceType(value, index) {
                this.baseinfo.serviceType = value.propertyValue
                this.showServiceTypePopView = false;
            },
            onConfirmArea(value, indexs) {
                this.areaSelected = value[0].name + value[1].name + value[2].name
                this.baseinfo.area = value[2].code
                this.showAreaPopView = false;
            },
            onConfirmSchoolSystem(value, index) {
                this.baseinfo.schoolSystem = value.propertyValue
                this.showSchoolSystemPopView = false;
            },
            onConfirmEducation(value, index) {
                this.baseinfo.highestEducation = value.propertyValue
                this.showEducationPopView = false;
            },

            goToPage(n) {
                this.$router.push({
                    'name': n
                });
            },
            getPropertyName(array, id) {
                let name = ""
                array.forEach(element => {
                    if (id == element.propertyValue) {
                        name = element.text;
                    }
                });
                return name
            },
            async initSelectOptions() {
                let result = await selectOptions();
                this.schoolColumns = result.data.dictGroup.school
                this.majorsStudiedColumns = result.data.dictGroup.study_major
                this.trainingFeeColumns = result.data.dictGroup.training_fee
                this.childAgeColumns = result.data.dictGroup.child_age
                this.incomeColumns = result.data.dictGroup.income
                this.serviceTypeColumns = result.data.dictGroup.service_type
                this.schoolSystemColumns = result.data.dictGroup.school_system
                this.paymentWayColumns = result.data.dictGroup.payment_way
                this.educationColumns = result.data.dictGroup.highest_education
                this.workingStateOption = result.data.dictGroup.working_state
                this.childTypeOption = result.data.dictGroup.child_type
                this.benefitsOption = result.data.dictGroup.benefits
                this.trainingNumberOption = result.data.dictGroup.training_number
                if (this.modify && this.modify == 1) {
                    // 获取基本信息
                    this.initUserBaseinfo();
                }
            },
            async initUserBaseinfo() {
                let result = await initUserBaseinfo()
                // 组装数据
                let obj = result.data.userBaseinfo
                if (obj.benefits == '0') {
                    obj.benefits = []
                } else {
                    obj.benefits = obj.benefits.split(',').map(Number)
                }
                if (obj.childType == '0') {
                    obj.childType = []
                } else {
                    obj.childType = obj.childType.split(',').map(Number)
                }

                obj.trainingNumber = obj.trainingNumber.split(',').map(Number)
                if (obj.area.length > 0) {
                    this.initAreas(obj.area)
                }
                if (!obj.trainingInfos) {
                    obj.trainingInfos = []
                }
                this.baseinfo = obj

            },
            async initAreas(id) {
                let result = await getArea(id);
                this.areaSelected = result.data.area
            },
            async saveBaseinfo() {
                let validated = true
                if (this.baseinfo.trainingNumber.length == 0) {
                    validated = false
                    Toast.fail("参加培训数量为必填项！")
                } else if (this.baseinfo.trainingFee == 0) {
                    validated = false
                    Toast.fail("我去年在培训上的花费为必填项！")
                } else if (this.trainingInfos.length > 0) {
                    this.trainingInfos.forEach((item, index) => {
                        if (item.trainingCourse == undefined || item.trainingCourse.length == 0) {
                            Toast.fail("培训" + (index + 1) + ": 培训课程不能为空！")
                            validated = false
                            return
                        } else if (item.beginTimeStr.length == 0) {
                            Toast.fail("培训" + (index + 1) + ": 开始日期不能为空！")
                            validated = false
                            return
                        } else if (item.endTimeStr.length == 0) {
                            Toast.fail("培训" + (index + 1) + ": 结束日期不能为空！")
                            validated = false
                            return
                        } else if (item.paymentWayStr.length == 0) {
                            Toast.fail("培训" + (index + 1) + ": 付费方式不能为空！")
                            validated = false
                            return
                        }
                    })
                }
                if (validated) {
                    let result = await saveBaseinfo(this.baseinfo)
                    Toast.success("保存成功！")
                    this.$router.push({path: "/home"});

                }
            }
        }
    }
</script>

<style lang="less" scoped>
    .van-nav-bar__title {
        font-size: 1rem;
    }
    .mainDiv {
        padding: 3.6rem 1rem 1rem 1rem;

        .title .van-cell__label {
            color: #108EE9;
            font-weight: bold;
            font-size: 0.8rem;
        }

        .px-title .van-cell__title {
            color: #8c939d;
        }

        .tips {
            color: #ea0b30;
            font-size: 0.65rem;
        }
    }

    /*.van-cell__title .van-field__label{
        font-size: 0.85rem;
    }
    .van-field__body{
        font-size: 0.8rem;
    }*/
    .buttonDiv {
        padding: 2rem 1rem 2rem 1rem;

        .button {
            margin: 0.3rem 0rem;
        }
    }
</style>
