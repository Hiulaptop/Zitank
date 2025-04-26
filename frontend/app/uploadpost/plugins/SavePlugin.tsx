
import { useLexicalComposerContext } from "@lexical/react/LexicalComposerContext";
import { COMMAND_PRIORITY_NORMAL, createCommand, LexicalCommand } from "lexical";
import { JSX, useEffect, useLayoutEffect } from "react";

export const SAVE_COMMAND: LexicalCommand<undefined> = createCommand('SAVE_COMMAND')

export default function SavePlugin(): JSX.Element | null {
    const [editor] = useLexicalComposerContext();
    useEffect(() => {
        return editor.registerCommand(SAVE_COMMAND, () => {
            console.log(JSON.stringify(editor.getEditorState().toJSON()))
            return true;
        }, COMMAND_PRIORITY_NORMAL)
    })
    useLayoutEffect(() => {
        const onKeyDown = (e: KeyboardEvent) => {
            if ((e.ctrlKey || e.metaKey) && e.key === 's') {
                e.preventDefault()
                editor.dispatchCommand(SAVE_COMMAND, undefined)
            }
        }
        return editor.registerRootListener(
            (
                rootElement: null | HTMLElement,
                prevRootElement: null | HTMLElement,
            ) => {
                if (prevRootElement !== null) {
                    prevRootElement.removeEventListener('keydown', onKeyDown);
                }
                if (rootElement !== null) {
                    rootElement.addEventListener('keydown', onKeyDown);
                }
            }
        );
    }, [editor])
    return null
}   