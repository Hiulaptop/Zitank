"use client"

import React from 'react';
import "./styles.css";
import { Editor } from './text-editor'
import { FlashMessageContext } from './context/FlashMessageContext';
// import { TypingAnimation } from './components/typing-animation';

export default function Home() {

  return (
    <div className="App">
      {/* <TypingAnimation /> */}
      <FlashMessageContext>
        <Editor />
      </FlashMessageContext>
    </div>
  )
}
