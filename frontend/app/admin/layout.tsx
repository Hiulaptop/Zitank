"use client";
import React from "react";
import Link from "next/link";
import { usePathname } from "next/navigation";

export default function AdminLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {

  const navValues = [
    { name: "home", value: "Tổng quan", src: "/admin" },
    { name: "post", value: "Quản lý bài viết", src: "/admin/posts" },
    { name: "rooms", value: "Quản lý phòng tập", src: "/admin/rooms" },
    { name: "orders", value: "Quản lý đơn hàng", src: "/admin/orders" },
    { name: "users", value: "Quản lý người dùng", src: "/admin/users" },
    { name: "finance", value: "Quản lý thanh toán & doanh thu", src: "/admin/financial" },
    { name: "musics", value: "Nhạc", src: "/admin/musics" },
    { name: "aboutus", value: "Cấu hình hệ thống", src: "/admin/advance" },
  ];

  const link = usePathname();
  const isLinkActive = (href: any) => {
    if (href === "/admin") {
      return link === href;
    }
    console.log(link, href);
    return link?.startsWith(href);
  }

  return (
    <div className="w-full h-full min-h-screen grid grid-cols-[15%_85%]">
      <div className=" bg-black pl-2">
        <nav className="flex flex-col mt-2 gap-1 text-lg font-bold text-white">
          {navValues.map((navValue) => (
            <Link key={navValue.name} href={navValue.src} className={`place-content-center text-center h-10 rounded-l-lg
                hover:bg-slate-200 hover:text-black
              ${isLinkActive(navValue.src) ? "bg-slate-100 text-black" : ""}`}>
              {navValue.value}
            </Link>
          ))}
        </nav>
      </div>
      <div className="">
        {children}
      </div>
    </div>
  );
}