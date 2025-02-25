import { loginAction } from './actions';
import Form from 'next/form';

export default function Login() {
  return (
    <div className="flex flex-col items-center h-full w-[500px] p-4 mx-auto rounded-lg border border-black">
      <h1 className="text-lg font-bold">Login</h1>
      <Form action={loginAction} className="grid grid-cols-[30%_70%] w-96 gap-2">
        
        <label htmlFor="username">Username:</label>
        <input className="rounded-md border border-black" type="text" id="username" name="username" required />
      

        <label htmlFor="password">Password:</label>
        <input className="rounded-md border border-black" type="password" id="password" name="password" required />
        

        <button type="submit" className="transition col-span-2 place-self-center w-36 h-8 rounded-md font-bold text-white bg-black hover:text-black hover:bg-white hover:ring-2 hover:ring-black duration-150">Login</button>
      </Form>
    </div>
  );
}