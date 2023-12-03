const std = @import("std");
const other = @import("part1/main.zig");
const fs = std.fs;
const heap = std.heap;

pub fn main() !void {
    const allocator = heap.page_allocator;
    var file = try fs.cwd().openFile("./src/input.txt", .{});
    defer file.close();

    const stat = try file.stat();
    const fileSize = stat.size;

    const a = try file.reader().readAllAlloc(allocator, fileSize);
    defer allocator.free(a);
    const sum = try other.solution(a);
    std.log.info("sum: {}\n", .{sum});
}
