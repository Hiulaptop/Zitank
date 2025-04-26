'use client';

import React, { useEffect } from 'react';
import { GetPosts } from './actions';

export default function AdminPage() {


    useEffect(() => {
        const data = GetPosts();
        data.then((v) => {
            if (v.status != "success")
                return;
            console.log(v.posts[0]);
        })
    }, []);

    return (
        <div className="flex flex-col w-full m-2">
            <div className='bg-red-300'>
                Quản lý bài viết
            </div>

            <div className='container w-full mx-auto grid grid-cols-[5%_70%_10%_10%_auto] gap-0.5 border-2 border-black bg-black'>
                <div className='h-12 text-center content-center text-lg font-bold bg-white'>STT</div>
                <div className='h-12 text-center content-center text-lg font-bold bg-white'>Tiêu đề</div>
                <div className='h-12 text-center content-center text-lg font-bold bg-white'>Ngày đăng</div>
                <div className='h-12 text-center content-center text-lg font-bold bg-white'>Thao tác</div>
                <div className='h-12 text-center content-center text-lg font-bold bg-white'></div>

                {/* test content */}
                <div className='h-8 text-center content-center text-lg font-bold bg-white'>1</div>
                <div className='h-8 text-center content-center text-lg font-bold bg-white'>Đây là test</div>
                <div className='h-8 text-center content-center text-lg font-bold bg-white'>03/04/2025</div>
                <div className='h-8 text-center content-center flex justify-center gap-2 py-1 bg-white'>
                    <button className='hover:text-blue-700'>Xóa</button>
                    <button className='hover:text-blue-700'>Sửa</button>
                </div>
                <div className='flex justify-center items-center h-8 bg-white'>
                    <input type="checkbox" className='w-5 h-5' />
                </div>


                <div className='h-8 text-center content-center text-lg font-bold bg-white'>2</div>
                <div className='h-8 text-center content-center text-lg font-bold bg-white'>Đây là test 2</div>
                <div className='h-8 text-center content-center text-lg font-bold bg-white'>03/04/2025</div>
                <div className='h-8 text-center content-center flex justify-center gap-2 py-1 bg-white'>
                    <button className='hover:text-blue-700'>Xóa</button>
                    <button className='hover:text-blue-700'>Sửa</button>
                </div>
                <div className='flex justify-center items-center h-8 bg-white'>
                    <input type="checkbox" className='w-5 h-5' />
                </div>
                {/* end test content */}


            </div>
        </div>
    );
}