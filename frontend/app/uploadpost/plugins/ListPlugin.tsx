import React from "react";
import { useLexicalComposerContext } from "@lexical/react/LexicalComposerContext";
import { 
    INSERT_ORDERED_LIST_COMMAND, 
    INSERT_UNORDERED_LIST_COMMAND,
    REMOVE_LIST_COMMAND,
} from "@lexical/list";
import { List, ListOrdered } from "lucide-react";

interface ListPluginProps{
    blockType: string;
}

export default function ListPlugin({ blockType }: ListPluginProps) {
    const [editor] = useLexicalComposerContext();

    return (
        <>
            <button
                onClick={() => {
                    if (blockType === 'ol') {
                        // blockType = 'paragraph';
                        editor.dispatchCommand(REMOVE_LIST_COMMAND, undefined);
                    }
                    else{
                        editor.dispatchCommand(INSERT_ORDERED_LIST_COMMAND, undefined);
                    }
                }}
                className={`w-10 h-10 rounded-xl disabled:cursor-not-allowed ${blockType === 'ol' ? "bg-blue-300" : "hover:bg-gray-300"}`}
            >
                <ListOrdered className="mx-auto"/>
            </button>

            <button
                onClick={() => {
                    if (blockType === 'ul') {
                        editor.dispatchCommand(REMOVE_LIST_COMMAND, undefined);
                    }
                    else{
                        editor.dispatchCommand(INSERT_UNORDERED_LIST_COMMAND, undefined);
                    }
                }}
                className={`w-10 h-10 rounded-xl disabled:cursor-not-allowed ${blockType === 'ul' ? "bg-blue-300" : "hover:bg-gray-300"}`}
            >
                <List className="mx-auto"/>
            </button>
        </>
    );
} 
