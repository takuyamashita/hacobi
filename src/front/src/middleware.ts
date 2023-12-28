import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'
 
export function middleware(request: NextRequest) {

  console.log(request.cookies.get('auth_token'));

  return NextResponse.next();
}
 
export const config = {
  matcher: '/livehouse/test/:path*',
}