#import <Cocoa/Cocoa.h>

#import "applescript.h"

bool execute(const char* s) {
    NSDictionary *error = [NSDictionary new];
    NSString *script = [NSString stringWithUTF8String:s];
    NSAppleScript *appleScript = [[NSAppleScript alloc] initWithSource:script];
    if ([appleScript executeAndReturnError:&error]) {
      return YES;
    } else {
      return NO;
    }
}