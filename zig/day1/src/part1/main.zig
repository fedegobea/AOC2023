const std = @import("std");
const ascii = @import("std").ascii;

const fs = std.fs;
const fmt = std.fmt;
const heap = std.heap;

fn get_last_number(str: []const u8) []const u8 {
    var result: []const u8 = "";
    const len = str.len;
    var i: usize = len;
    while (i > 0) : (i -= 1) {
        var a = i - 1;
        result = str[a..i];
        const is_digit = ascii.isDigit(str[a]);
        if (is_digit) {
            break;
        }
    }
    return result;
}

fn get_first_number(str: []const u8) []const u8 {
    var result: []const u8 = "";
    for (str, 0..) |char, i| {
        result = str[i .. i + 1];
        const is_digit = ascii.isDigit(char);
        if (is_digit) {
            break;
        }
    }
    return result;
}

fn get_first_and_last_number(str: []const u8) !i32 {
    const first = get_first_number(str);
    const last = get_last_number(str);
    const allocator = heap.page_allocator;
    const first_and_last = try std.fmt.allocPrint(allocator, "{s}{s}", .{ first, last });
    defer allocator.free(first_and_last);

    return try std.fmt.parseInt(i32, first_and_last, 10);
}

pub fn solution(input: []u8) !i32 {
    var it = std.mem.split(u8, input, "\n");
    var sum: i32 = 0;
    while (it.next()) |token| {
        const res = get_first_and_last_number(token);
        sum += try res;
    }

    return sum;
}

const expect = std.testing.expect;

test "part1" {
    var file = try fs.cwd().openFile("./src/part1/testinput.txt", .{});
    defer file.close();
    const stat = try file.stat();
    const fileSize = stat.size;
    const allocator = heap.page_allocator;

    const a = try file.reader().readAllAlloc(allocator, fileSize);
    defer allocator.free(a);
    const sol: i32 = try solution(a);
    try expect(sol == 142);
}
