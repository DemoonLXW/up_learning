import { useState } from "react";
import styled, { keyframes } from "styled-components";

const SlideshowContainer = styled.div`
        position: relative;
    `;

const fade = keyframes`
    from {
        opacity: 0.4;
    }
    to {
        opacity: 1;
    }

`;

const ImageContainer = styled.div`
    display: none;
    animation: ${fade} 1.2s;

    &.show {
        display: block;
    }

    & > img {
        width: 100%;
    }
`;

const SelectDots = styled.div`
    text-align: center;
    position: absolute;
    bottom: 10px;
    width: 100%;
    & > span {
        display: inline-block;
        width: 15px;
        height: 15px;
        border-radius: 50%;
        background-color: #ffffff;
        user-select: none;
        cursor: pointer;
        opacity: 0.5;
        margin: 8px;
    }

    & > span:hover, & > span.selected {
        opacity: 1;
    }

`;



export default function Slideshow(props: {urls?: string[]}) {
    
    if(!props.urls || props.urls.length<=0) {
        return (<div></div>);
    }

    const [selected, setSelected] = useState(0);
    const handleClick = (index: number) => {
        setSelected(index);
    }
    
    const images = props.urls.map((url, index)=> <ImageContainer key={index} className={selected == index ? "show" : ""} ><img src={url} /></ImageContainer>);
    const dots = props.urls.map((_, index)=> <span key={index}  className={selected == index ? "selected" : ""} onClick={() => handleClick(index)} ></span>);


    return (
        <SlideshowContainer>
            {images}
            <SelectDots>
                {dots}
            </SelectDots>
        </SlideshowContainer>
    );


}