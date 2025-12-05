// import type { ReactNode } from "react";

import { useEffect, useRef, type Ref } from "react";


// type PropsWithChildren<P> = P & { children?: ReactNode };

function RevealAnimation({children}: {children?: React.ReactNode}) {
    const ref:Ref<HTMLDivElement> = useRef(null);

    useEffect(() => {
        const observer = new IntersectionObserver(
            ([entry]) => {
                if(entry.isIntersecting) {
                    ref.current?.classList.add('visible');
                }
            },{ threshold: 0.2, rootMargin: '0px 0px -50px 0px'});

        if(ref.current){
            observer.observe(ref.current);
        }

        return () => observer.disconnect();
    });

    return (
    <div ref={ref} className="reveal flex flex-col h-full">
        {children}
    </div>);
}

export default RevealAnimation;