import type { Route } from "./+types/home";
import { Calculator } from "~/calculator/Calculator";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "Calculator" },
    { name: "description", content: "Welcome to the calculator!" },
  ];
}

export default function Home() {
  return <Calculator />;
}
