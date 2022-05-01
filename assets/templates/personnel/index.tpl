<!DOCTYPE html>
<html>
<head>
    <title>行为溯源</title>
    <meta charset="UTF-8">
    <link type="text/css" rel="stylesheet" href="../../bootstarp/css/trace.css"/>
    <link rel="stylesheet" href="/static/element-ui/lib/theme-chalk/index.css?ver=1">
</head>
<body>
<div id="content">
    <div id="main" v-cloak>
        <!--        <el-button @click="drawer = true" type="primary" style="margin-left: 16px;">-->
        <!--            点我打开-->
        <!--        </el-button>-->
        <div style="margin:10px 5px">
            <span>筛选方式:</span>
            <el-select v-model="time_mode" placeholder="请选择" size="small" @change="getData">
                <el-option
                        v-for="item in time_options"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value">
                </el-option>
            </el-select>

            <el-date-picker
                    v-show="time_mode === 2 "
                    size="small"
                    v-model="date_value"
                    type="datetimerange"
                    :picker-options="pickerOptions"
                    range-separator="至"
                    start-placeholder="开始日期"
                    end-placeholder="结束日期"
                    align="right">
            </el-date-picker>
        </div>
        <div class="content-flex flex-justify-content-space-between">
            <div id="traceChar">
                <div id="mountNode"></div>
            </div>
            <div id="nodeInfo" class="nodeInfo">
                <div>本地</div>
                <div class="infoMainNav">
                    <el-descriptions class="margin-top" title="节点信息" :column="1" :size="size" border>
                        <template slot="extra">
<!--                            <el-button v-show="local_terminalInfo.show" type="primary" size="small" plain-->
<!--                                       @click="local_terminalInfo.show=!local_terminalInfo.show">收起-->
<!--                            </el-button>-->
<!--                            <el-button v-show="!local_terminalInfo.show" type="primary" size="small" plain-->
<!--                                       @click="local_terminalInfo.show=!local_terminalInfo.show">展开-->
<!--                            </el-button>-->
                        </template>

                        <el-descriptions-item label-class-name="test" v-for="item,index in local_terInfoL" v-if="item.show==true">
                            <template slot="label">
                                [[item.label]]
                            </template>
                            [[item.value]]
                        </el-descriptions-item>

                    </el-descriptions>

                    <el-descriptions class="margin-top" title="行为信息" :column="1" :size="size" border>
                        <template slot="extra">
<!--                            <el-button type="primary" size="small" plain>操作</el-button>-->
                        </template>

                        <el-descriptions-item label-class-name="test" v-for="item,index in local_attrInfo" v-if="item.show==true">
                            <template slot="label">
                                [[item.label]]
                            </template>
                            [[item.value]]
                        </el-descriptions-item>
                    </el-descriptions>


                    <!--                <div class="infoMainNavTitle">本地-->
                    <!--                    <el-button size="mini" style="float:right" type="success"-->
                    <!--                               @click="dialogOptionalClientTableVisible=true">处理-->
                    <!--                    </el-button>-->
                    <!--                </div>-->
                    <!--                <div class="terminalInfo">-->
                    <!--                    <p>节点信息</p>-->
                    <!--                    <ul>-->
                    <!--                        <li><span class="cont">发生时间: </span><span class="contValue">[[local_terminalInfo.time]]</span>-->
                    <!--                        </li>-->
                    <!--                        <li><span class="cont">主机名称:</span><span class="contValue">[[local_terminalInfo.name]]</span>-->
                    <!--                        </li>-->
                    <!--                    </ul>-->
                    <!--                </div>-->
                    <!--                <div class="attackInfo">-->
                    <!--                    <p>行为信息</p>-->
                    <!--                    <ul>-->
                    <!--                        <li><span class="cont">PUID:</span><span class="contValue">[[local_attackInfo.puid]]</span></li>-->
                    <!--                        <li><span class="cont">MD5:</span><span class="contValue">[[local_attackInfo.md5]]</span></li>-->
                    <!--                        <li><span class="cont">严重程度:</span><span class="contValue">[[local_attackInfo.lavel]]</span>-->
                    <!--                        </li>-->
                    <!--                        &lt;!&ndash;                        <li><span class="cont">处理方法&手段:</span><span class="contValue">[[local_attackInfo.method]]</span></li>&ndash;&gt;-->
                    <!--                        <li><span class="cont">行为:</span><span class="contValue">[[local_attackInfo.behavior]]</span>-->
                    <!--                        </li>-->
                    <!--                        <li><span class="cont">文件路径:</span><span class="contValue">[[local_attackInfo.filePath]]</span>-->
                    <!--                        </li>-->
                    <!--                        <li><span class="cont">文件来源:</span><span class="contValue">[[local_attackInfo.source]]</span>-->
                    <!--                        </li>-->
                    <!--                        <li><span class="cont">运行时间:</span><span class="contValue">[[local_attackInfo.startTime]]至[[local_attackInfo.endTime]]</span>-->
                    <!--                        </li>-->
                    <!--                        <li><span class="cont">持续时间:</span><span class="contValue">[[local_attackInfo.duration]]</span>-->
                    <!--                        </li>-->
                    <!--                    </ul>-->
                    <!--                </div>-->
                </div>

            </div>
        </div>

        <div id="dialog" v-cloak>
            <el-drawer
                    title="我是标题"
                    :visible.sync="drawer"
                    :with-header="false">
                <span>我来啦!</span>
            </el-drawer>
        </div>
    </div>
</div>
<script src="https://gw.alipayobjects.com/os/antv/pkg/_antv.g6-3.7.1/dist/g6.min.js"></script>
<script src="static/js/LAB.min.js"></script>
<script src="static/js/cfg/cfg_trace.js"></script>

</body>

</html>