"use client";

import { Carousel, CarouselContent, CarouselItem, CarouselNext, CarouselPrevious } from "@/components/ui/carousel";

export function Banner() {
    const data = [
        {
            "id": 1,
            "image": "/bannerimg/test1.jpg",
        },
        {
            "id": 2,
            "image": "/bannerimg/test2.jpg",
        },
        {
            "id": 3,
            "image": "/bannerimg/test3.jpg",
        },
        {
            "id": 4,
            "image": "/bannerimg/test4.jpg",
        },
    ];

    return (
        <Carousel opts={{ loop: true }} className="container mx-auto relative flex items-center">
            <CarouselContent>
                {data.map((items, id) => (
                    <CarouselItem key={id}>
                        <img src={items.image} />
                    </CarouselItem>
                ))}

            </CarouselContent>
            <CarouselPrevious className="absolute inset-y-0 left-0 px-2 h-12 rounded-xl bg-gray-300" />
            <CarouselNext className="absolute inset-y-0 right-0 px-2 h-12 rounded-xl bg-gray-300" />
        </Carousel >
    );
}