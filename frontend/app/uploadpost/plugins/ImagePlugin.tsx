import React, { useRef, useState } from "react";
import { useLexicalComposerContext } from "@lexical/react/LexicalComposerContext";
import { Image } from "lucide-react";
import { $createImageNode } from "../nodes/ImageNode";
import { $insertNodes } from "lexical";

export default function ImagePlugin() {
    const [editor] = useLexicalComposerContext();
    const [file, setFile] = useState<File | null>(null);
    const inputRef = useRef<HTMLInputElement>(null);

    const onAddImage = (file: File) => {
        const src = URL.createObjectURL(file);
        editor.update(() => {
            const node = $createImageNode({ src, altText: "Dummy text" });
            $insertNodes([node]);
        });
    };

    return (
        <div>
            <button
                onClick={() => inputRef?.current?.click()}
                className="w-10 h-10 rounded-xl disabled:cursor-not-allowed hover:bg-gray-300"
            >
                <Image className="mx-auto" />
            </button>
            <input
                type="file"
                ref={inputRef}
                accept="image/*"
                style={{ display: "none" }}
                onChange={(e) => {
                    const file = e.target.files?.[0];
                    if (file) {
                        setFile(file);
                        onAddImage(file);
                    }
                }}
            />
        </div>
    );
}
