//定义默认状态
let initState = {
    todos: [
        {
            id: parseInt(Math.random() * 10000000),
            isComplete: false,
            title: '学习redux'
        }, {
            id: parseInt(Math.random() * 10000000),
            isComplete: true,
            title: '学习react'
        }, {
            id: parseInt(Math.random() * 10000000),
            isComplete: false,
            title: '学习node'
        }
    ]
};
export default initState;