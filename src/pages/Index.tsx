import Slideshow from '../components/Slideshow'
import img1 from '../assets/images/img_lights_wide.jpg'
import img2 from '../assets/images/img_mountains_wide.jpg'
import img3 from '../assets/images/img_nature_wide.jpg'
import img4 from '../assets/images/img_snow_wide.jpg'

const urls: string[] = [
    img1, img2, img3, img4
];

export default () => (
    <div>
        <h1>首页</h1>   
        <Slideshow urls={urls} />
    </div>
);

