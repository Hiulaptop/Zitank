import { useLexicalComposerContext } from "@lexical/react/LexicalComposerContext";
import { mergeRegister, $getNearestNodeOfType } from '@lexical/utils';
import {
    $getSelection,
    $isRangeSelection,
    CAN_REDO_COMMAND,
    CAN_UNDO_COMMAND,
    FORMAT_ELEMENT_COMMAND,
    FORMAT_TEXT_COMMAND,
    REDO_COMMAND,
    SELECTION_CHANGE_COMMAND,
    UNDO_COMMAND
} from "lexical";
import { AlignCenter, AlignJustify, AlignLeft, AlignRight, Bold, Italic, PartyPopper, Redo, Strikethrough, Underline, Undo } from "lucide-react";
import React, { useCallback, useEffect, useState } from "react";
import { $isListNode, ListNode } from "@lexical/list";


import ColorPlugin from "./ColorPlugin";
import ListPlugin from "./ListPlugin";
import ImagePlugin from "./ImagePlugin";
import YoutubePlugin from "./YoutubePlugin";

export default function Toolbars() {
    const [editor] = useLexicalComposerContext();

    const [canUndo, setCanUndo] = useState(false);
    const [canRedo, setCanRedo] = useState(false);
    const [isBold, setIsBold] = useState(false);
    const [isItalic, setIsItalic] = useState(false);
    const [isUnderline, setIsUnderline] = useState(false);
    const [isStrikethrough, setIsStrikethrough] = useState(false);
    const [blockType, setBlockType] = useState('paragraph');

    const $updateToolbar = useCallback(() => {
        const selection = $getSelection();
        if ($isRangeSelection(selection)) {
            setIsBold(selection.hasFormat('bold'));
            setIsItalic(selection.hasFormat('italic'));
            setIsUnderline(selection.hasFormat('underline'));
            setIsStrikethrough(selection.hasFormat('strikethrough'));
            const anchorNode = selection.anchor.getNode();
            const element = anchorNode.getKey() === 'root' ? anchorNode : anchorNode.getTopLevelElementOrThrow();
            const elementKey = element.getKey();
            const elementDom = editor.getElementByKey(elementKey);
            
            if (!elementDom) return;

            if ($isListNode(element)){
                const parentList = $getNearestNodeOfType(anchorNode, ListNode);
                const type = parentList ? parentList.getTag() : element.getTag();
                setBlockType(type);
            }
        };



    }, []);

    useEffect(() => {
        return mergeRegister(
            editor.registerUpdateListener(({ editorState }) => {
                editorState.read(() => {
                    $updateToolbar();
                });
            }),
            editor.registerCommand(
                SELECTION_CHANGE_COMMAND,
                (_payload, _newEditor) => {
                    $updateToolbar();
                    return false;
                },
                1,
            ),
            editor.registerCommand(
                CAN_UNDO_COMMAND,
                (payload) => {
                    setCanUndo(payload);
                    return false;
                },
                1,
            ),
            editor.registerCommand(
                CAN_REDO_COMMAND,
                (payload) => {
                    setCanRedo(payload);
                    return false;
                },
                1,
            ),
        );
    }, [editor, $updateToolbar]);


    return (
        <div id="toolbar" className="flex flex-row h-12 py-1 gap-1 border-b-1 rounded-t-xl bg-gray-100">
            <button
                id="Undo"
                disabled={!canUndo}
                onClick={() => {
                    editor.dispatchCommand(UNDO_COMMAND, undefined)
                }}
                className={`w-10 rounded-xl ${canUndo && "hover:bg-gray-300"} disabled:cursor-not-allowed`}
            >
                <Undo className={`mx-auto ${!canUndo ? "text-gray-500" : "text-black"}`} />
            </button>
            <button
                id="Redo"
                disabled={!canRedo}
                onClick={() => {
                    editor.dispatchCommand(REDO_COMMAND, undefined)
                }}
                className={`w-10 rounded-xl ${canRedo && "hover:bg-gray-300"} disabled:cursor-not-allowed`}
            >
                <Redo className={`mx-auto ${!canRedo ? "text-gray-500" : "text-black"}`} />
            </button>

            <button
                id="Bold"
                onClick={() => {
                    editor.dispatchCommand(FORMAT_TEXT_COMMAND, 'bold');
                }}
                className={`w-10 rounded-xl hover:bg-gray-300 ${isBold && "bg-gray-300"}`}
            >
                <Bold className={`mx-auto`} />
            </button>

            <button
                id="Italic"
                onClick={() => {
                    editor.dispatchCommand(FORMAT_TEXT_COMMAND, 'italic');
                }}
                className={`w-10 rounded-xl hover:bg-gray-300 ${isItalic && "bg-gray-300"}`}
            >
                <Italic className={`mx-auto`} />
            </button>

            <button
                id="Underline"
                onClick={() => {
                    editor.dispatchCommand(FORMAT_TEXT_COMMAND, 'underline');
                }}
                className={`w-10 rounded-xl hover:bg-gray-300 ${isUnderline && "bg-gray-300"}`}
            >
                <Underline className={`mx-auto`} />
            </button>

            <button
                id="Strikethrough"
                onClick={() => {
                    editor.dispatchCommand(FORMAT_TEXT_COMMAND, 'strikethrough');
                }}
                className={`w-10 rounded-xl hover:bg-gray-300 ${isStrikethrough && "bg-gray-300"}`}
            >
                <Strikethrough className={`mx-auto`} />
            </button>

            <button
                onClick={() => {
                    editor.dispatchCommand(FORMAT_ELEMENT_COMMAND, 'left');
                }}
                className={`w-10 rounded-xl hover:bg-gray-300`}
            >
                <AlignLeft className={`mx-auto`} />

            </button>
            
            <button
                onClick={() => {
                    editor.dispatchCommand(FORMAT_ELEMENT_COMMAND, 'right');
                }}
                className={`w-10 rounded-xl hover:bg-gray-300 `}
            >
                <AlignRight className={`mx-auto`}/>
            </button>
            
            <button
                onClick={() => {
                    editor.dispatchCommand(FORMAT_ELEMENT_COMMAND, 'center');
                }}
                className={`w-10 rounded-xl hover:bg-gray-300 `}
            >
                <AlignCenter className={`mx-auto`}/>
            </button>
            
            
            <button
                onClick={() => {
                    editor.dispatchCommand(FORMAT_ELEMENT_COMMAND, 'justify');
                }}
                className={`w-10 rounded-xl hover:bg-gray-300 `}
            >
                <AlignJustify className={`mx-auto`}/>
            </button>
            
            <ColorPlugin />
            <ListPlugin blockType = {blockType} />
            <ImagePlugin />
            <YoutubePlugin />
        </div>
    );
}

