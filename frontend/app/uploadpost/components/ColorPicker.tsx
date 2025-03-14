
import React, { useEffect, useRef, useState } from "react";
import { PhotoshopPicker, SketchPicker } from 'react-color';

interface ColorPickerProps {
    color: string;
    onChange: (color: string) => void;
    Icon: React.ReactNode;
}

export default function ColorPicker({ color, onChange, Icon }: ColorPickerProps) {
    const [isOpen, setIsOpen] = useState(false);
    const ref = useRef<HTMLDivElement>(null);

    useEffect(() => {
        const handleOutSideClick = (event: any) => {
            if (!ref.current?.contains(event.target)) {
                // alert("Outside Clicked.");
                // console.log("Outside Clicked. ");
                setIsOpen(false);
            }
        };
        if(ref){
            window.addEventListener("mousedown", handleOutSideClick);
            return () => {
                window.removeEventListener("mousedown", handleOutSideClick);
            };
        }
    }, [ref]);

    return (
        <div ref={ref} className={`relative w-10 h-10`}>
            <button
                onClick={() => setIsOpen(!isOpen)}
                className={`w-10 h-10 rounded-xl ${isOpen && "bg-gray-300"} hover:bg-gray-300 disabled:cursor-not-allowed`}
            >
                {Icon}
            </button>
            {isOpen &&
                <SketchPicker
                    color={color}
                    onChange={(color) => onChange(color.hex)}
                    className="absolute top-8 z-10 select-none"
                />
            }
        </div>
    );
}
