import React, { useEffect, useMemo, useState } from "react";
import ColorPicker from "../components/ColorPicker";
import { Type, PaintBucket } from "lucide-react";
import { useLexicalComposerContext } from "@lexical/react/LexicalComposerContext";
import { $getSelection, $isRangeSelection, BaseSelection, SELECTION_CHANGE_COMMAND } from "lexical";
import { $getSelectionStyleValueForProperty, $patchStyleText } from "@lexical/selection";
import { mergeRegister } from '@lexical/utils';


export default function ColorPlugin() {
    const [editor] = useLexicalComposerContext();
    const [{ color, background }, setColor] = useState({ color: 'black', background: 'white' });

    useMemo(() => {
        editor.update(() => {
            const selection = $getSelection();
            if (selection) {
                $patchStyleText(selection, { color: color, background: background });
            }
        })
    }, [color, background])
    const $updateToolbar = () => {
        const select = $getSelection();
        if ($isRangeSelection(select)) {
            const color = $getSelectionStyleValueForProperty(select, "color", "#000000");
            const background = $getSelectionStyleValueForProperty(select, "background", "#ffffff");
            setColor({ color, background });
        }
    }

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
        );
    }, [editor]);

    return (
        <>
            <ColorPicker
                color={color}
                onChange={(color: string) => {
                    setColor({ color, background });
                }}
                Icon={
                    <Type color={color} className="mx-auto" />
                }
            />

            <ColorPicker
                color={background}
                onChange={(background: string) => {
                    setColor({ color, background });
                }}
                Icon={
                    <PaintBucket className="mx-auto" />
                }
            />
        </>
    );
} 
