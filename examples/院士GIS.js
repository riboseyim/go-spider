var data = [
{name:'江苏',value:124},
{name:'上海',value:86},
{name:'浙江',value:81},
{name:'安徽',value:44},
{name:'福建',value:44},
{name:'湖南',value:42},
{name:'山东',value:38},
{name:'北京',value:37},
{name:'广东',value:31},
{name:'四川',value:29},
{name:'湖北',value:26},
{name:'辽宁',value:24},
{name:'河北',value:22},
{name:'陕西',value:21},
{name:'江西',value:20},
{name:'河南',value:19},
{name:'重庆',value:18},
{name:'吉林',value:14},
{name:'香港',value:13},
{name:'天津',value:12},
{name:'黑龙江',value:10},
{name:'山西',value:10},
{name:'甘肃',value:9},
{name:'云南',value:7},
{name:'广西',value:5},
{name:'贵州',value:5},
{name:'内蒙古',value:4},
{name:'海南',value:1},
{name:'宁夏',value:1},
{name:'无资料',value:1},
{name:'新疆',value:1}
];
var geoCoordMap = {
    '北京':[116.46,39.92],
    '天津':[117.2,39.13],
    '上海':[121.48,31.22],
    '重庆':[106.54,29.59],
    '西藏':[91.11,29.97],
    '新疆':[87.68,43.77],
    '福建':[119.3,26.08],
    '广东':[113.23,23.16],
    '山西':[112.53,37.87],
    '云南':[102.73,25.04],
    '海南':[110.35,20.02],
    '辽宁':[121.62,38.92],
    '沈阳':[123.38,41.8],
    '吉林':[125.35,43.88],
    '江西':[115.89,28.68],
    '内蒙古':[111.65,40.82],
    '四川':[104.06,30.67],
    '广西':[110.28,25.29],
    '江苏':[118.78,32.04],
    '贵州':[106.71,26.57],
    '浙江':[120.19,30.26],
    '山东':[117,36.65],
    '甘肃':[103.73,36.03],
    '河南':[113.65,34.76],
    '黑龙江':[126.63,45.75],
    '湖南':[113,28.21],
    '安徽':[117.27,31.86],
    '湖北':[114.31,30.52],
    '宁夏':[106.27,38.47]
};

var convertData = function (data) {
    var res = [];
    for (var i = 0; i < data.length; i++) {
        var geoCoord = geoCoordMap[data[i].name];
        if (geoCoord) {
            res.push({
                name: data[i].name,
                value: geoCoord.concat(data[i].value)
            });
        }
    }
    return res;
};

option = {
    title: {
        text: '中国科学院院士分布情况',
        subtext: '数据来源：中科院官网',
        sublink: 'https://riboseyim.github.io',
        left: 'center'
    },
    tooltip : {
        trigger: 'item'
    },
    bmap: {
        center: [104.114129, 37.550339],
        zoom: 5,
        roam: true,
        mapStyle: {
            styleJson: [{
                'featureType': 'water',
                'elementType': 'all',
                'stylers': {
                    'color': '#d1d1d1'
                }
            }, {
                'featureType': 'land',
                'elementType': 'all',
                'stylers': {
                    'color': '#f3f3f3'
                }
            }, {
                'featureType': 'railway',
                'elementType': 'all',
                'stylers': {
                    'visibility': 'off'
                }
            }, {
                'featureType': 'highway',
                'elementType': 'all',
                'stylers': {
                    'color': '#fdfdfd'
                }
            }, {
                'featureType': 'highway',
                'elementType': 'labels',
                'stylers': {
                    'visibility': 'off'
                }
            }, {
                'featureType': 'arterial',
                'elementType': 'geometry',
                'stylers': {
                    'color': '#fefefe'
                }
            }, {
                'featureType': 'arterial',
                'elementType': 'geometry.fill',
                'stylers': {
                    'color': '#fefefe'
                }
            }, {
                'featureType': 'poi',
                'elementType': 'all',
                'stylers': {
                    'visibility': 'off'
                }
            }, {
                'featureType': 'green',
                'elementType': 'all',
                'stylers': {
                    'visibility': 'off'
                }
            }, {
                'featureType': 'subway',
                'elementType': 'all',
                'stylers': {
                    'visibility': 'off'
                }
            }, {
                'featureType': 'manmade',
                'elementType': 'all',
                'stylers': {
                    'color': '#d1d1d1'
                }
            }, {
                'featureType': 'local',
                'elementType': 'all',
                'stylers': {
                    'color': '#d1d1d1'
                }
            }, {
                'featureType': 'arterial',
                'elementType': 'labels',
                'stylers': {
                    'visibility': 'off'
                }
            }, {
                'featureType': 'boundary',
                'elementType': 'all',
                'stylers': {
                    'color': '#fefefe'
                }
            }, {
                'featureType': 'building',
                'elementType': 'all',
                'stylers': {
                    'color': '#d1d1d1'
                }
            }, {
                'featureType': 'label',
                'elementType': 'labels.text.fill',
                'stylers': {
                    'color': '#999999'
                }
            }]
        }
    },
    series : [
        {
            name: 'pm2.5',
            type: 'scatter',
            coordinateSystem: 'bmap',
            data: convertData(data),
            symbolSize: function (val) {
                return val[2] / 10;
            },
            label: {
                normal: {
                    formatter: '{b}',
                    position: 'right',
                    show: false
                },
                emphasis: {
                    show: true
                }
            },
            itemStyle: {
                normal: {
                    color: 'purple'
                }
            }
        },
        {
            name: 'Top 5',
            type: 'effectScatter',
            coordinateSystem: 'bmap',
            data: convertData(data.sort(function (a, b) {
                return b.value - a.value;
            }).slice(0, 6)),
            symbolSize: function (val) {
                return val[2] / 10;
            },
            showEffectOn: 'render',
            rippleEffect: {
                brushType: 'stroke'
            },
            hoverAnimation: true,
            label: {
                normal: {
                    formatter: '{b}',
                    position: 'right',
                    show: true
                }
            },
            itemStyle: {
                normal: {
                    color: 'purple',
                    shadowBlur: 10,
                    shadowColor: '#333'
                }
            },
            zlevel: 1
        }
    ]
};
