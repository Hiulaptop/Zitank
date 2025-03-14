"use client";

import { $getRoot, $getSelection } from 'lexical';
import { useEffect, useState } from 'react';

import { AutoFocusPlugin } from '@lexical/react/LexicalAutoFocusPlugin';
import { LexicalComposer } from '@lexical/react/LexicalComposer';
import { RichTextPlugin } from '@lexical/react/LexicalRichTextPlugin';
import { ContentEditable } from '@lexical/react/LexicalContentEditable';
import { HistoryPlugin } from '@lexical/react/LexicalHistoryPlugin';
import { ListPlugin } from '@lexical/react/LexicalListPlugin';
import { LexicalErrorBoundary } from '@lexical/react/LexicalErrorBoundary';
import { useLexicalComposerContext } from '@lexical/react/LexicalComposerContext';

import {ListNode, ListItemNode} from "@lexical/list"

import { theme } from "./Theme";
import Toolbars from './ToolbarsPlugin';

import { ImageNode } from "../nodes/ImageNode";
import { YoutubeNode } from "../nodes/YoutubeNode";

function onError(error: any) {
  console.error(error);
}

export default function Editor() {
  const initialConfig = {
    namespace: 'PostEditor',
    theme: theme,
    onError,
    nodes: [ListNode, 
      ListItemNode,
      ImageNode,
      YoutubeNode,
    ],
  };

  const [editorState, setEditorState] = useState<any>();
  function onChangeState(editorState: any) {
    const editorStateJSON = editorState.toJSON();
    setEditorState(JSON.stringify(editorStateJSON));

    // console.log(editorStateJSON);
  }

  function MyOnChangePlugin(onChange: any) {
    const [editor] = useLexicalComposerContext();
    useEffect(() => {
      return editor.registerUpdateListener(({ editorState }) => {
        onChangeState(editorState);
      });
    }, [editor, onChange]);
    return null;
  }

  return (
    <div className='relative min-h-[500px] w-4xl mx-auto bg-white shadow-2xl rounded-lg'>
      <LexicalComposer initialConfig={initialConfig}>
        <Toolbars />
        <RichTextPlugin
          contentEditable={
            <ContentEditable
              placeholder={<></>}
              className='relative min-h-[500px] hover:outline-none focus:outline-none px-2'
              aria-placeholder={'Enter some text...'}
            />
          }
          ErrorBoundary={LexicalErrorBoundary}
        />
        <ListPlugin />
        <HistoryPlugin />
        <AutoFocusPlugin />
        <MyOnChangePlugin onChange={onChangeState} />
      </LexicalComposer>
    </div>
  );
}