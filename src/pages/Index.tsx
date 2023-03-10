import Slideshow from '../components/Slideshow'
import img1 from '../assets/images/img_lights_wide.jpg'
import img2 from '../assets/images/img_mountains_wide.jpg'
import img3 from '../assets/images/img_nature_wide.jpg'
import img4 from '../assets/images/img_snow_wide.jpg'
import TopBar from '../components/TopBar'


const urls: string[] = [
    img1, img2, img3, img4
];

export default function Index() {
    return (
        <div>
            <TopBar />
            <Slideshow urls={urls} />
        </div>
    );
}
