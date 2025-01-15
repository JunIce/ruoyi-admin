import deleteIcon from '../assets/icons/delete.svg';
import clearIcon from '../assets/icons/clear.svg';
import saveIcon from '../assets/icons/save.svg';
import printerIcon from '../assets/icons/printer.svg';
import resetZoomIcon from '../assets/icons/reset-zoom.svg';
import zoomInIcon from '../assets/icons/zoomin.svg';
import zoomOutIcon from '../assets/icons/zoomout.svg';


export default [
    {
        name: '放大',
        icon: zoomInIcon,
        command: 'zoomIn',
    },
    {
        name: '缩小',
        icon: zoomOutIcon,
        command: 'zoomOut',
    },
    {
        name: '重置缩放',
        icon: resetZoomIcon,
        command: 'resetZoom',
    },
    {
        name: '清空',
        icon: clearIcon,
        command: 'clear',
    },
    // {
    //     name: '打印',
    //     icon: printerIcon,
    //     command: 'print',
    // },
    {
        name: '删除',
        icon: deleteIcon,
        command: 'delete',
    },
    // {
    //     name: '保存',
    //     icon: saveIcon,
    //     command: 'save',
    // },
]