"use client";

import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import Logo from "../public/logo/logo.svg";
import FooterImg from "../public/logo/logo2.svg";
import Image from "next/image";
import Link from "next/link";
import { useEffect, useRef, useState } from "react";
import { usePathname } from "next/navigation";
import Cookies from "js-cookie";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

// export const metadata: Metadata = {
//   title: "Create Next App",
//   description: "Generated by create next app",
// };

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {

  const navValues = [
    { name: "home", value: "Trang chủ", src: "/" },
    { name: "post", value: "Bài viết", src: "/posts" },
    { name: "rooms", value: "Đặt phòng", src: "/rooms" },
    { name: "musics", value: "Nhạc", src: "/musics" },
    { name: "aboutus", value: "Về chúng tôi", src: "/aboutus" },
  ];

  const [menuOpen, setMenuOpen] = useState(false);
  const [isLogin, setIsLogin] = useState(false);
  const [navItems, setNavItems] = useState("home");
  
  const loginRef = useRef<HTMLDivElement>(null);
  const logoutRef = useRef<HTMLDivElement>(null);
  useEffect(()=>{
    if (loginRef.current) {
      loginRef.current.style.display = isLogin ? "none" : "flex";
    }
    if (logoutRef.current) {
      logoutRef.current.style.display = isLogin ? "flex" : "none";
    }
  }, [isLogin]);

  const link = usePathname();
  const isLinkActive = (href: any) => {
    if (href === "/") {
      return link === href;
    }
    return link?.startsWith(href);
  }

  useEffect(() => {
    // const check = () => {
      const token = Cookies.get("jwt");
      setIsLogin(!!token);
    // }

    // check();
  }, [link]);

  return (
    <html lang="vi">
      <body className={`${geistSans.variable} ${geistMono.variable} antialiased`}>
        <header className="fixed flex flex-row items-center justify-between top-0 h-24 w-full md:px-[20%] bg-black text-white z-10">
          <Link href="/" className="w-auto h-full">
            <Image src={Logo} alt="logo" className="w-auto h-full" />
          </Link>

          <nav className="hidden h-full md:visible md:grid grid-cols-5 gap-1 text-lg font-bold">
            {navValues.map((navValue) => (
              <Link key={navValue.name} href={navValue.src} className={`place-content-center text-center ${isLinkActive(navValue.src)
                      ? "text-blue-600"
                      : "text-white"}`}>	
                {navValue.value}
              </Link>
            ))}
          </nav>

          <div className="hidden md:flex flex-row gap-4" ref={loginRef}>
            <Link href="/login">Login</Link>
            <Link href="/signup">Signup</Link>
          </div>

          <div className="hidden md:flex flex-row gap-4" ref={logoutRef}>
            <Link href="/profile" >Profile</Link>
            <Link href="/logout" >Logout</Link>
          </div>

          <div className="md:hidden px-4 relative">
            <button className="flex items-center justify-center w-10 h-10  focus:outline-none" onClick={() => setMenuOpen(!menuOpen)}>
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                <line x1="4" y1="6" x2="20" y2="6" />
                <line x1="4" y1="12" x2="20" y2="12" />
                <line x1="4" y1="18" x2="20" y2="18" />
              </svg>
            </button>
            {menuOpen &&
              (
                <div className="absolute right-0 z-10 mt-2 w-56 origin-top-right rounded-md bg-white ring-1 shadow-lg ring-black/5 focus:outline-none">
                  <nav className="" role="none">
                    {navValues.map((navValue) => (
                      <Link key={navValue.name} href={navValue.src} className={`block px-4 py-2 text-sm hover:bg-slate-200 text-black ${isLinkActive(navValue.src)
                          ? "text-blue-600"
                          : "text-black"}`}>{navValue.value}
                        </Link>
                    ))};
                  </nav>
                </div>
              )
            }
          </div>
        </header>


        <main className="mt-24 h-auto min-h-screen">
          {children}
        </main>


        <footer className="h-auto min-h-96 w-full pt-8 text-white bg-black">
          <div className="container flex md:flex-row flex-col w-full h-full mx-auto">
            {/* need to change to svg */}
            <div className="flex flex-col gap-4 h-full w-[500px]" id="image-zone">
              <Image src={FooterImg} alt="logo" className="max-w-72" />
              <p>
                slogan
              </p>
              <nav className="grid grid-cols-8" role="none">
                <Link href="/post" className="mx-auto">
                  <svg xmlns="http://www.w3.org/2000/svg" className="transition hover:scale-110" x="0px" y="0px" width="32" height="32" viewBox="0 0 50 50" fill="white">
                    <path d="M25,3C12.85,3,3,12.85,3,25c0,11.03,8.125,20.137,18.712,21.728V30.831h-5.443v-5.783h5.443v-3.848 c0-6.371,3.104-9.168,8.399-9.168c2.536,0,3.877,0.188,4.512,0.274v5.048h-3.612c-2.248,0-3.033,2.131-3.033,4.533v3.161h6.588 l-0.894,5.783h-5.694v15.944C38.716,45.318,47,36.137,47,25C47,12.85,37.15,3,25,3z"></path>
                  </svg>
                </Link>

                <Link href="/rooms" className="mx-auto">
                  <svg xmlns="http://www.w3.org/2000/svg" className="transition hover:scale-110" x="0px" y="0px" width="32" height="32" viewBox="0 0 50 50" fill="white">
                    <path d="M 16 3 C 8.83 3 3 8.83 3 16 L 3 34 C 3 41.17 8.83 47 16 47 L 34 47 C 41.17 47 47 41.17 47 34 L 47 16 C 47 8.83 41.17 3 34 3 L 16 3 z M 37 11 C 38.1 11 39 11.9 39 13 C 39 14.1 38.1 15 37 15 C 35.9 15 35 14.1 35 13 C 35 11.9 35.9 11 37 11 z M 25 14 C 31.07 14 36 18.93 36 25 C 36 31.07 31.07 36 25 36 C 18.93 36 14 31.07 14 25 C 14 18.93 18.93 14 25 14 z M 25 16 C 20.04 16 16 20.04 16 25 C 16 29.96 20.04 34 25 34 C 29.96 34 34 29.96 34 25 C 34 20.04 29.96 16 25 16 z"></path>
                  </svg>
                </Link>

                {/* <Link href="/musics" className="mx-auto">
                  <svg xmlns="http://www.w3.org/2000/svg" className="transition hover:scale-110" x="0px" y="0px" width="32" height="32" viewBox="0 0 50 50" fill="white">
                    <path d="M 11 4 C 7.134 4 4 7.134 4 11 L 4 39 C 4 42.866 7.134 46 11 46 L 39 46 C 42.866 46 46 42.866 46 39 L 46 11 C 46 7.134 42.866 4 39 4 L 11 4 z M 13.085938 13 L 21.023438 13 L 26.660156 21.009766 L 33.5 13 L 36 13 L 27.789062 22.613281 L 37.914062 37 L 29.978516 37 L 23.4375 27.707031 L 15.5 37 L 13 37 L 22.308594 26.103516 L 13.085938 13 z M 16.914062 15 L 31.021484 35 L 34.085938 35 L 19.978516 15 L 16.914062 15 z"></path>
                  </svg>
                </Link> */}
                <Link href="/user" className="mx-auto">
                  <svg xmlns="http://www.w3.org/2000/svg" className="transition hover:scale-110" x="0px" y="0px" width="32" height="32" viewBox="0 0 50 50" fill="white">
                    <path d="M 44.898438 14.5 C 44.5 12.300781 42.601563 10.699219 40.398438 10.199219 C 37.101563 9.5 31 9 24.398438 9 C 17.800781 9 11.601563 9.5 8.300781 10.199219 C 6.101563 10.699219 4.199219 12.199219 3.800781 14.5 C 3.398438 17 3 20.5 3 25 C 3 29.5 3.398438 33 3.898438 35.5 C 4.300781 37.699219 6.199219 39.300781 8.398438 39.800781 C 11.898438 40.5 17.898438 41 24.5 41 C 31.101563 41 37.101563 40.5 40.601563 39.800781 C 42.800781 39.300781 44.699219 37.800781 45.101563 35.5 C 45.5 33 46 29.398438 46.101563 25 C 45.898438 20.5 45.398438 17 44.898438 14.5 Z M 19 32 L 19 18 L 31.199219 25 Z"></path>
                  </svg>
                </Link>
              </nav>
            </div>


          </div>

          {/* Copyright */}
          <div className="w-full h-12 max-h-32 text-center content-end mt-full">
            Copyright &copy; WEBSITE URL | Design & Development by
            &nbsp;
            <Link href="https://github.com/Hiulaptop" className="hover:text-blue-600">
              Trung Hieu
            </Link>
            &nbsp;
            and
            &nbsp;
            <Link href="https://github.com/NgTHung" className="hover:text-blue-600">
              Tuan Hung
            </Link>
          </div>
        </footer>
      </body>
    </html>
  );
}
