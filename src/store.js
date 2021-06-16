import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

// Div A
/*
let fieldWidth = 1800;
let fieldLength = 2200;
let penAreaWidth = 500;
let penAreaDepth = 150;
let goalWidth = 400;
let goalDepth = 150;
let centerCircleRadius = 200;
*/

// Div B

let fieldWidth = 1300;
let fieldLength = 1500;
let penAreaWidth = 700;
let penAreaDepth = 150;
let goalWidth = 400;
let goalDepth = 100;
let centerCircleRadius = 200;

// Internal
let robotSize = 80;
let triangleSize = 70;


let defaultField = {
    activeSourceId: '',
    sources: {},
    fieldWidth: fieldWidth,
    fieldLength: fieldLength,
    boundaryWidth: 150,
    penAreaWidth: penAreaWidth,
    penAreaDepth: penAreaDepth,
    goalWidth: goalWidth,
    goalDepth: goalDepth,
    robotSize: robotSize,
    triangleSize: triangleSize,
    centerCircleRadius: centerCircleRadius,
    ballRadius: 21.5,
    shapes: [
        // left 
        {
            line: {
                p1: {x: -fieldLength / 2, y: -fieldWidth / 2 + triangleSize},
                p2: {x: -fieldLength / 2, y: fieldWidth / 2 - triangleSize}
            },
        },
        // left-bottom
        {
            line: {
                p1: {x: -fieldLength / 2, y: fieldWidth / 2 - triangleSize},
                p2: {x: -fieldLength / 2 + triangleSize, y: fieldWidth / 2}
            }
        },
        // bottom
        {
            line: {
                p1: {x: -fieldLength / 2 + triangleSize, y: fieldWidth / 2},
                p2: {x: fieldLength / 2 - triangleSize, y: fieldWidth / 2}
            }
        },
        // bottom-right
        {
            line: {
                p1: {x: fieldLength / 2 - triangleSize, y: fieldWidth / 2},
                p2: {x: fieldLength / 2, y: fieldWidth / 2 - triangleSize}
            }
        },
        // right
        {
            line: {
                p1: {x: fieldLength / 2, y: fieldWidth / 2 - triangleSize},
                p2: {x: fieldLength / 2, y: -fieldWidth / 2 + triangleSize}
            }
        },
        // right-top
        {
            line: {
                p1: {x: fieldLength / 2, y: -fieldWidth / 2 + triangleSize},
                p2: {x: fieldLength / 2 - triangleSize, y: -fieldWidth / 2}
            }
        },
        // top
        {
            line: {
                p1: {x: fieldLength / 2 - triangleSize, y: -fieldWidth / 2},
                p2: {x: -fieldLength / 2 + triangleSize, y: -fieldWidth / 2}
            }
        },   
        // top-left
        {
            line: {
                p1: {x: -fieldLength / 2 + triangleSize, y: -fieldWidth / 2},
                p2: {x: -fieldLength / 2, y: -fieldWidth / 2 + triangleSize}
            }
        },   
        // left penalty area
        {
            rect: {
                x: -fieldLength / 2,
                y: -penAreaWidth / 2,
                width: penAreaDepth,
                height: penAreaWidth,
                fill: '',
                fillOpacity: 0
            }
        },  
        // right penalty area
        {
            rect: {
                x: fieldLength / 2 - penAreaDepth,
                y: -penAreaWidth / 2,
                width: penAreaDepth,
                height: penAreaWidth,
                fill: '',
                fillOpacity: 0
            }
        },
        // mid line
        {
            line: {
                p1: {x: 0, y: -fieldWidth / 2},
                p2: {x: 0, y: fieldWidth / 2}
            }
        },
        // left goal
        {
            rect: {
                x: -fieldLength / 2 - goalDepth,
                y: -goalWidth / 2,
                width: goalDepth,
                height: goalWidth,
                fill: '',
                fillOpacity: 0
            }
        },
        // right goal
        {
            rect: {
                x: fieldLength / 2,
                y: -goalWidth / 2,
                width: goalDepth,
                height: goalWidth,
                fill: '',
                fillOpacity: 0
            }
        },
        // center circle
        {
            circle: {
                center: {x: 0, y: 0},
                radius: centerCircleRadius,
                stroke: 'white',
                strokeWidth: 10,
                fill: '',
                fillOpacity: 0
            }
        },
        // robot rect [test]
        {
            rect: {
                x: 0 - robotSize/2,
                y: 0 - robotSize/2,
                width: robotSize,
                height: robotSize,
                ori: 0,
                stroke: 'black',
                strokeWidth: '10',
                fill: 'yellow',
                fillOpacity: 1
            }
        },
        // robot text [test]
        {
            text: {
                text: '1',
                p: {x: 0, y: 0},
                fill: 'black',
            }
        }
    ],
};

export default new Vuex.Store({
    state: {
        field: defaultField
    },
    mutations: {
        SOCKET_ONOPEN() {
        },
        SOCKET_ONCLOSE() {
        },
        SOCKET_ONERROR() {
        },
        SOCKET_ONMESSAGE(state, message) {
            state.field = message;
        },
        SOCKET_RECONNECT() {
        },
        SOCKET_RECONNECT_ERROR() {
        },
    }
});
