import styled from "styled-components";


const TopBarContainer = styled.div`
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100vw;
    height: 5.5rem;
`;

const Wrapper = styled.div`
    position: relative;
    width: 80%;
    height: 100%;
`;

const Logo = styled.div`
    position: absolute;
    top: 0;
    left: 0;
    height: 100%;
    display: flex;
    align-items: center;
    & > div {
        font-size: 2.5rem;
        line-height: 2.5rem;
    }

    @media only screen and (max-width: 768px) {
        position: relative;
    }
`;

const Options = styled.div`
    position: absolute;
    top: 0;
    right: 0;
    display: flex;
    align-items: center;
    justify-content: space-around;
    height: 100%;
    & > div {
        margin-left: calc(70/19.2*1vw);
        font-size: 2rem;
        line-height: 2rem;
        user-select: none;
        cursor: pointer;
    }

    @media only screen and (max-width: 768px) {
        position: relative;
        justify-content: start;
        
        & > div {
            margin-right: calc(70/19.2*1vw);
            margin-left: 0;
        }
    }

`;

export default function TopBar() {
    

    return (
        <TopBarContainer>
            <Wrapper>
                <Logo><div>Up Learning</div></Logo>
                <Options>
                    <div>注册</div>
                    <div>登录</div>
                    <div>语言</div>
                    <div>关于</div>
                </Options>
            </Wrapper>
        </TopBarContainer>
    );
}